# Code Smell Burndown

Tracked issues from a review of the five command-line tools under `cmd/`.

## Refactoring

- [x] **Heavy duplication between `mysql2sqlite-data` and `mysql2sqlite-verify`.**
  `config`, `loadSetupDefaults`, `validateConfig`, `mysqlDSN`, `resolveTables`,
  `sqliteTablesInCreateOrder`, `countRows`, `mysqlIdent`, `sqliteIdent`, and `fail`
  are copy-pasted across both tools (~80 lines). Extract a shared package
  (e.g. `internal/mybbdb`).

- [x] **Regex compiled on every call in `mysql2sqlite`.**
  `parseColumn` (line 150) calls `regexp.MustCompile` inside the function body,
  recompiling on every column. Hoist to a package-level `var`.

- [x] **Raw SQL string threading in `countRows`.**
  `countRows` accepts a fully-formed SQL string built with `Sprintf`. The
  signature invites misuse. Accept `(db, table)` and build the query internally.

## Error Handling

- [x] **Ignored write errors in `mysql2sqlite`.**
  `writeTable` discards errors from `fmt.Fprintf`/`fmt.Fprintln`, and
  `defer w.Flush()` in `main` discards the flush error. A write failure is
  silently swallowed.

- [x] **No context timeout in `-data` and `-verify`.**
  Both use a bare `context.Background()`. A hung MySQL connection or large
  import blocks forever. Add a `--timeout` flag or a default deadline.

## Consistency

- [x] **Inconsistent error/exit style across tools.**
  `mysql2sqlite`, `-data`, and `-verify` use a hand-rolled `fail()` that writes
  to stderr via `fmt.Fprintln` and calls `os.Exit(1)`. `gobbs-serve` and
  `gobbs-static` use `log.Fatal`. Pick one convention — `log.Fatal` is
  idiomatic and adds a timestamp.

## Hygiene

- [x] **`defer tx.Rollback()` after possible commit in `mysql2sqlite-data`.**
  The deferred `Rollback` fires after a successful `Commit` (returns
  `sql.ErrTxDone`). Harmless but noisy; guard the defer.

## Tests to Add

Unit tests for pure functions that currently have zero coverage. These should
be written before the refactoring items above so the extractions are
regression-safe.

- [x] **`cmd/mysql2sqlite`: `mapType`** — verify MySQL-to-SQLite type mapping
  for each branch (`tinyint`, `varchar`, `blob`, `bool`, unknown, etc.).
- [x] **`cmd/mysql2sqlite`: `parseColumn`** — verify DEFAULT extraction,
  NOT NULL detection, and AUTO_INCREMENT detection from raw column tails.
- [x] **`cmd/mysql2sqlite`: `normalizeDefault`** — verify BLOB `""` → `X''`
  conversion and pass-through for other types.
- [x] **`cmd/mysql2sqlite`: `parseColumns`** — verify backtick stripping and
  whitespace trimming on comma-separated column lists.
- [ ] **`cmd/mysql2sqlite-data`: `coerceSQLiteValue`** — verify nil, `[]byte`,
  `time.Time`, `bool`, `int64`/`float64`/`string`, and fallback coercions.
- [x] **Shared helpers (`mysqlIdent`, `sqliteIdent`)** — verify identifier
  quoting and escaping of embedded quotes.
