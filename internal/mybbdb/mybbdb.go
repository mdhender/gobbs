// Package mybbdb provides shared configuration, connection helpers, and
// SQL utilities used by the mysql2sqlite-data and mysql2sqlite-verify tools.
package mybbdb

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mdhender/gobbs/internal/setupjson"
)

// Config holds the flags shared by the data-import and verify tools.
type Config struct {
	MysqlAddr     string
	MysqlDatabase string
	MysqlUser     string
	MysqlPassword string
	SQLitePath    string
	SetupFile     string
	AllTables     bool
	Timeout       time.Duration
}

// RegisterFlags binds the standard flag set to cfg's fields.
func RegisterFlags(cfg *Config) {
	flag.StringVar(&cfg.SetupFile, "setup-file", "setup.json", "path to setup.json")
	flag.StringVar(&cfg.MysqlAddr, "mysql-addr", "", "MySQL host:port")
	flag.StringVar(&cfg.MysqlDatabase, "mysql-db", "", "MySQL database name")
	flag.StringVar(&cfg.MysqlUser, "mysql-user", "", "MySQL username")
	flag.StringVar(&cfg.MysqlPassword, "mysql-password", "", "MySQL password")
	flag.StringVar(&cfg.SQLitePath, "sqlite-path", "mybb.sqlite3", "SQLite database path")
	flag.BoolVar(&cfg.AllTables, "all-tables", false, "process all tables in target schema order")
	flag.DurationVar(&cfg.Timeout, "timeout", 0, "overall operation timeout (e.g. 5m, 1h); 0 means no timeout")
}

// ContextWithTimeout returns a context with the configured timeout applied.
// If Timeout is zero, the parent context is returned unchanged with a no-op
// cancel function.
func ContextWithTimeout(parent context.Context, cfg Config) (context.Context, context.CancelFunc) {
	if cfg.Timeout > 0 {
		return context.WithTimeout(parent, cfg.Timeout)
	}
	return parent, func() {}
}

// LoadSetupDefaults fills in any unset Config fields from setup.json. If the
// file does not exist, no error is returned.
func LoadSetupDefaults(cfg *Config) error {
	values, err := setupjson.Parse(cfg.SetupFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if cfg.MysqlAddr == "" {
		cfg.MysqlAddr = values.Database.Hostname
	}
	if cfg.MysqlDatabase == "" {
		cfg.MysqlDatabase = values.Database.Database
	}
	if cfg.MysqlUser == "" {
		cfg.MysqlUser = values.Database.Username
	}
	if cfg.MysqlPassword == "" {
		cfg.MysqlPassword = values.Database.Password
	}
	return nil
}

// ValidateConfig returns an error if required fields are missing.
func ValidateConfig(cfg Config) error {
	if cfg.MysqlAddr == "" {
		return errors.New("missing MySQL host:port; set --mysql-addr or setup.json database.hostname")
	}
	if cfg.MysqlDatabase == "" {
		return errors.New("missing MySQL database name; set --mysql-db or setup.json database.database")
	}
	if cfg.MysqlUser == "" {
		return errors.New("missing MySQL username; set --mysql-user or setup.json database.username")
	}
	return nil
}

// MysqlDSN builds a go-sql-driver/mysql DSN from cfg.
func MysqlDSN(cfg Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC",
		cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlDatabase)
}

// ResolveTables returns the list of tables to process. When cfg.AllTables is
// set it reads the SQLite schema; otherwise it uses the positional args.
func ResolveTables(ctx context.Context, sqliteDB *sql.DB, cfg Config, args []string) ([]string, error) {
	if cfg.AllTables {
		tables, err := SQLiteTablesInCreateOrder(ctx, sqliteDB)
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

// SQLiteTablesInCreateOrder returns all non-internal tables ordered by rowid
// (i.e. creation order).
func SQLiteTablesInCreateOrder(ctx context.Context, db *sql.DB) ([]string, error) {
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

// CountRows counts the rows in the given table. The quoteFn parameter should
// be MysqlIdent or SQLiteIdent depending on the database.
func CountRows(ctx context.Context, db *sql.DB, table string, quoteFn func(string) string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", quoteFn(table))
	var count int64
	if err := db.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// MysqlIdent quotes a MySQL identifier with backticks.
func MysqlIdent(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

// SQLiteIdent quotes a SQLite identifier with double-quotes.
func SQLiteIdent(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}

// Fail prints err to stderr and exits with code 1.
func Fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
