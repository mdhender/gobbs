package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mdhender/gobbs/internal/mybbdb"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

func main() {
	cfg := mybbdb.Config{}
	mybbdb.RegisterFlags(&cfg)
	flag.Parse()

	if err := mybbdb.LoadSetupDefaults(&cfg); err != nil {
		mybbdb.Fail(err)
	}
	if err := mybbdb.ValidateConfig(cfg); err != nil {
		mybbdb.Fail(err)
	}

	ctx := context.Background()

	mysqlDB, err := sql.Open("mysql", mybbdb.MysqlDSN(cfg))
	if err != nil {
		mybbdb.Fail(err)
	}
	defer mysqlDB.Close()

	sqliteDB, err := sql.Open("sqlite", cfg.SQLitePath)
	if err != nil {
		mybbdb.Fail(err)
	}
	defer sqliteDB.Close()

	if err := mysqlDB.PingContext(ctx); err != nil {
		mybbdb.Fail(fmt.Errorf("connect mysql: %w", err))
	}
	if err := sqliteDB.PingContext(ctx); err != nil {
		mybbdb.Fail(fmt.Errorf("connect sqlite: %w", err))
	}
	if _, err := sqliteDB.ExecContext(ctx, "PRAGMA foreign_keys = OFF"); err != nil {
		mybbdb.Fail(fmt.Errorf("disable sqlite foreign keys: %w", err))
	}

	tables, err := mybbdb.ResolveTables(ctx, sqliteDB, cfg, flag.Args())
	if err != nil {
		mybbdb.Fail(err)
	}

	var failed []string
	for _, table := range tables {
		if err := importTable(ctx, mysqlDB, sqliteDB, table); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			failed = append(failed, table)
		}
	}
	if len(failed) > 0 {
		mybbdb.Fail(fmt.Errorf("import completed with issues in %d table(s): %s", len(failed), strings.Join(failed, ", ")))
	}
}

func importTable(ctx context.Context, mysqlDB, sqliteDB *sql.DB, table string) error {
	declaredTypes, err := sqliteDeclaredTypes(ctx, sqliteDB, table)
	if err != nil {
		return err
	}
	sourceCount, err := mybbdb.CountRows(ctx, mysqlDB, table, mybbdb.MysqlIdent)
	if err != nil {
		return fmt.Errorf("%s: count mysql rows: %w", table, err)
	}

	rows, err := mysqlDB.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s", mybbdb.MysqlIdent(table)))
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

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s", mybbdb.SQLiteIdent(table))); err != nil {
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
			args[i] = coerceSQLiteValue(values[i], declaredTypes[col])
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

	targetCount, err := mybbdb.CountRows(ctx, sqliteDB, table, mybbdb.SQLiteIdent)
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
	rows, err := db.QueryContext(ctx, fmt.Sprintf("PRAGMA table_info(%s)", mybbdb.SQLiteIdent(table)))
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

func buildInsertSQL(table string, columns []string) string {
	colNames := make([]string, 0, len(columns))
	placeholders := make([]string, 0, len(columns))
	for _, col := range columns {
		colNames = append(colNames, mybbdb.SQLiteIdent(col))
		placeholders = append(placeholders, "?")
	}
	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		mybbdb.SQLiteIdent(table),
		strings.Join(colNames, ", "),
		strings.Join(placeholders, ", "),
	)
}

// coerceSQLiteValue adapts driver values to SQLite-compatible types without
// altering forum content semantics.
func coerceSQLiteValue(v any, declType string) any {
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
