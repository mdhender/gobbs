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
	flag.BoolVar(&cfg.allTables, "all-tables", false, "verify all tables in target schema order")
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

	tables, err := resolveTables(ctx, sqliteDB, cfg, flag.Args())
	if err != nil {
		fail(err)
	}

	mismatches := 0
	for _, table := range tables {
		ok, err := verifyTable(ctx, mysqlDB, sqliteDB, table)
		if err != nil {
			fail(err)
		}
		if !ok {
			mismatches++
		}
	}

	if mismatches > 0 {
		fail(fmt.Errorf("verification failed for %d table(s)", mismatches))
	}
	fmt.Printf("verified %d table(s): all row counts match\n", len(tables))
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

func verifyTable(ctx context.Context, mysqlDB, sqliteDB *sql.DB, table string) (bool, error) {
	mysqlCount, err := countRows(ctx, mysqlDB, fmt.Sprintf("SELECT COUNT(*) FROM %s", mysqlIdent(table)))
	if err != nil {
		return false, fmt.Errorf("%s: count mysql rows: %w", table, err)
	}
	sqliteCount, err := countRows(ctx, sqliteDB, fmt.Sprintf("SELECT COUNT(*) FROM %s", sqliteIdent(table)))
	if err != nil {
		return false, fmt.Errorf("%s: count sqlite rows: %w", table, err)
	}

	status := "OK"
	match := mysqlCount == sqliteCount
	if !match {
		status = "MISMATCH"
	}
	fmt.Printf("%-30s mysql=%-8d sqlite=%-8d %s\n", table, mysqlCount, sqliteCount, status)
	return match, nil
}

func countRows(ctx context.Context, db *sql.DB, query string) (int64, error) {
	var count int64
	if err := db.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
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
