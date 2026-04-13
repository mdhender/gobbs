package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

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

	tables, err := mybbdb.ResolveTables(ctx, sqliteDB, cfg, flag.Args())
	if err != nil {
		mybbdb.Fail(err)
	}

	mismatches := 0
	for _, table := range tables {
		ok, err := verifyTable(ctx, mysqlDB, sqliteDB, table)
		if err != nil {
			mybbdb.Fail(err)
		}
		if !ok {
			mismatches++
		}
	}

	if mismatches > 0 {
		mybbdb.Fail(fmt.Errorf("verification failed for %d table(s)", mismatches))
	}
	fmt.Printf("verified %d table(s): all row counts match\n", len(tables))
}

func verifyTable(ctx context.Context, mysqlDB, sqliteDB *sql.DB, table string) (bool, error) {
	mysqlCount, err := mybbdb.CountRows(ctx, mysqlDB, fmt.Sprintf("SELECT COUNT(*) FROM %s", mybbdb.MysqlIdent(table)))
	if err != nil {
		return false, fmt.Errorf("%s: count mysql rows: %w", table, err)
	}
	sqliteCount, err := mybbdb.CountRows(ctx, sqliteDB, fmt.Sprintf("SELECT COUNT(*) FROM %s", mybbdb.SQLiteIdent(table)))
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
