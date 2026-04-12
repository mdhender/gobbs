package main

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

type config struct {
	mysqlAddr     string
	mysqlDatabase string
	mysqlUser     string
	mysqlPassword string
	sqlitePath    string
	envFile       string
	allTables     bool
}

func main() {
	cfg := config{}
	flag.StringVar(&cfg.envFile, "env-file", ".env", "path to .env file")
	flag.StringVar(&cfg.mysqlAddr, "mysql-addr", "127.0.0.1:3307", "MySQL host:port")
	flag.StringVar(&cfg.mysqlDatabase, "mysql-db", "", "MySQL database name")
	flag.StringVar(&cfg.mysqlUser, "mysql-user", "", "MySQL username")
	flag.StringVar(&cfg.mysqlPassword, "mysql-password", "", "MySQL password")
	flag.StringVar(&cfg.sqlitePath, "sqlite-path", "mybb.sqlite3", "SQLite database path")
	flag.BoolVar(&cfg.allTables, "all-tables", false, "import all tables in target schema order")
	flag.Parse()

	if err := loadEnvDefaults(&cfg); err != nil {
		fail(err)
	}
	if err := validateConfig(cfg); err != nil {
		fail(err)
	}

	ctx := context.Background()

	mysqlDB, err := sql.Open("mysql", mysqlDSN(cfg))
	if err != nil {
		fail(err)
	}
	defer mysqlDB.Close()

	sqliteDB, err := sql.Open("sqlite", cfg.sqlitePath)
	if err != nil {
		fail(err)
	}
	defer sqliteDB.Close()

	if err := mysqlDB.PingContext(ctx); err != nil {
		fail(fmt.Errorf("connect mysql: %w", err))
	}
	if err := sqliteDB.PingContext(ctx); err != nil {
		fail(fmt.Errorf("connect sqlite: %w", err))
	}
	if _, err := sqliteDB.ExecContext(ctx, "PRAGMA foreign_keys = OFF"); err != nil {
		fail(fmt.Errorf("disable sqlite foreign keys: %w", err))
	}

	tables, err := resolveTables(ctx, sqliteDB, cfg, flag.Args())
	if err != nil {
		fail(err)
	}

	for _, table := range tables {
		if err := importTable(ctx, mysqlDB, sqliteDB, table); err != nil {
			fail(err)
		}
	}
}

func loadEnvDefaults(cfg *config) error {
	values, err := parseDotEnv(cfg.envFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if cfg.mysqlDatabase == "" {
		cfg.mysqlDatabase = values["MYBB_DATABASE_DATABASE"]
	}
	if cfg.mysqlUser == "" {
		cfg.mysqlUser = values["MYBB_DATABASE_USERNAME"]
	}
	if cfg.mysqlPassword == "" {
		cfg.mysqlPassword = values["MYBB_DATABASE_PASSWORD"]
	}
	return nil
}

func validateConfig(cfg config) error {
	if cfg.mysqlDatabase == "" {
		return errors.New("missing MySQL database name; set --mysql-db or MYBB_DATABASE_DATABASE")
	}
	if cfg.mysqlUser == "" {
		return errors.New("missing MySQL username; set --mysql-user or MYBB_DATABASE_USERNAME")
	}
	return nil
}

func mysqlDSN(cfg config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC", cfg.mysqlUser, cfg.mysqlPassword, cfg.mysqlAddr, cfg.mysqlDatabase)
}

func parseDotEnv(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	values := map[string]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return nil, fmt.Errorf("invalid line in %s: %s", path, line)
		}
		values[strings.TrimSpace(key)] = trimQuotes(strings.TrimSpace(value))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return values, nil
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func resolveTables(ctx context.Context, sqliteDB *sql.DB, cfg config, args []string) ([]string, error) {
	if cfg.allTables {
		tables, err := sqliteTablesInCreateOrder(ctx, sqliteDB)
		if err != nil {
			return nil, err
		}
		if len(args) > 0 {
			return nil, errors.New("do not pass table names with --all-tables")
		}
		return tables, nil
	}
	if len(args) == 0 {
		return nil, errors.New("pass one or more table names, or use --all-tables")
	}
	return args, nil
}

func sqliteTablesInCreateOrder(ctx context.Context, db *sql.DB) ([]string, error) {
	rows, err := db.QueryContext(ctx, `
SELECT name
FROM sqlite_master
WHERE type = 'table' AND name NOT LIKE 'sqlite_%'
ORDER BY rowid`)
	if err != nil {
		return nil, fmt.Errorf("list sqlite tables: %w", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, name)
	}
	return tables, rows.Err()
}

func importTable(ctx context.Context, mysqlDB, sqliteDB *sql.DB, table string) error {
	declaredTypes, err := sqliteDeclaredTypes(ctx, sqliteDB, table)
	if err != nil {
		return err
	}
	sourceCount, err := countRows(ctx, mysqlDB, fmt.Sprintf("SELECT COUNT(*) FROM %s", mysqlIdent(table)))
	if err != nil {
		return fmt.Errorf("%s: count mysql rows: %w", table, err)
	}

	rows, err := mysqlDB.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s", mysqlIdent(table)))
	if err != nil {
		return fmt.Errorf("%s: query mysql rows: %w", table, err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("%s: list columns: %w", table, err)
	}

	insertSQL := buildInsertSQL(table, columns)
	tx, err := sqliteDB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: begin sqlite transaction: %w", table, err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s", sqliteIdent(table))); err != nil {
		return fmt.Errorf("%s: clear sqlite table: %w", table, err)
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM sqlite_sequence WHERE name = ?", table); err != nil {
		return fmt.Errorf("%s: clear sqlite sequence: %w", table, err)
	}

	stmt, err := tx.PrepareContext(ctx, insertSQL)
	if err != nil {
		return fmt.Errorf("%s: prepare sqlite insert: %w", table, err)
	}
	defer stmt.Close()

	values := make([]any, len(columns))
	dest := make([]any, len(columns))
	for i := range values {
		dest[i] = &values[i]
	}

	inserted := int64(0)
	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			return fmt.Errorf("%s: scan mysql row: %w", table, err)
		}
		args := make([]any, len(columns))
		for i, col := range columns {
			args[i] = normalizeValue(values[i], declaredTypes[col])
		}
		if _, err := stmt.ExecContext(ctx, args...); err != nil {
			return fmt.Errorf("%s: insert sqlite row: %w", table, err)
		}
		inserted++
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("%s: iterate mysql rows: %w", table, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit sqlite transaction: %w", table, err)
	}

	targetCount, err := countRows(ctx, sqliteDB, fmt.Sprintf("SELECT COUNT(*) FROM %s", sqliteIdent(table)))
	if err != nil {
		return fmt.Errorf("%s: count sqlite rows: %w", table, err)
	}
	if sourceCount != targetCount {
		return fmt.Errorf("%s: row count mismatch mysql=%d sqlite=%d", table, sourceCount, targetCount)
	}

	fmt.Printf("%s: imported %d rows\n", table, inserted)
	return nil
}

func sqliteDeclaredTypes(ctx context.Context, db *sql.DB, table string) (map[string]string, error) {
	rows, err := db.QueryContext(ctx, fmt.Sprintf("PRAGMA table_info(%s)", sqliteIdent(table)))
	if err != nil {
		return nil, fmt.Errorf("%s: read sqlite table info: %w", table, err)
	}
	defer rows.Close()

	types := map[string]string{}
	for rows.Next() {
		var (
			cid        int
			name       string
			declType   string
			notNull    int
			defaultVal sql.NullString
			pk         int
		)
		if err := rows.Scan(&cid, &name, &declType, &notNull, &defaultVal, &pk); err != nil {
			return nil, fmt.Errorf("%s: scan sqlite table info: %w", table, err)
		}
		types[name] = strings.ToUpper(declType)
	}
	if len(types) == 0 {
		return nil, fmt.Errorf("%s: table not found in sqlite", table)
	}
	return types, rows.Err()
}

func countRows(ctx context.Context, db *sql.DB, query string) (int64, error) {
	var count int64
	if err := db.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func buildInsertSQL(table string, columns []string) string {
	colNames := make([]string, 0, len(columns))
	placeholders := make([]string, 0, len(columns))
	for _, col := range columns {
		colNames = append(colNames, sqliteIdent(col))
		placeholders = append(placeholders, "?")
	}
	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		sqliteIdent(table),
		strings.Join(colNames, ", "),
		strings.Join(placeholders, ", "),
	)
}

func normalizeValue(v any, declType string) any {
	switch value := v.(type) {
	case nil:
		return nil
	case []byte:
		if strings.Contains(declType, "BLOB") {
			copyBuf := make([]byte, len(value))
			copy(copyBuf, value)
			return copyBuf
		}
		return string(value)
	case time.Time:
		return value.UTC().Format(time.RFC3339Nano)
	case bool:
		if value {
			return 1
		}
		return 0
	case int64, float64, string:
		return value
	default:
		return fmt.Sprint(value)
	}
}

func mysqlIdent(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

func sqliteIdent(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
