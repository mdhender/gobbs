package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type columnDef struct {
	name          string
	rawType       string
	mappedType    string
	notNull       bool
	defaultValue  string
	autoIncrement bool
}

type indexDef struct {
	name     string
	columns  []string
	unique   bool
	fulltext bool
}

type tableDef struct {
	name       string
	columns    []columnDef
	primaryKey []string
	indexes    []indexDef
}

var (
	createTableRE = regexp.MustCompile("^CREATE TABLE `([^`]+)` \\($")
	columnRE      = regexp.MustCompile("^`([^`]+)`\\s+(.+)$")
	indexRE       = regexp.MustCompile("^(UNIQUE\\s+)?INDEX `([^`]+)` \\((.+)\\)$")
	fulltextRE    = regexp.MustCompile("^FULLTEXT INDEX `([^`]+)` \\((.+)\\)$")
	pkRE          = regexp.MustCompile("^PRIMARY KEY \\((.+)\\)$")
)

func main() {
	inPath := flag.String("in", "mysql-schema.sql", "input MySQL schema file")
	outPath := flag.String("out", "sqlite-schema.sql", "output SQLite schema file")
	flag.Parse()

	tables, err := parseTables(*inPath)
	if err != nil {
		fail(err)
	}

	out, err := os.Create(*outPath)
	if err != nil {
		fail(err)
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	defer w.Flush()

	fmt.Fprintln(w, "-- Generated from mysql-schema.sql by mysql2sqlite.go")
	fmt.Fprintln(w, "PRAGMA foreign_keys = OFF;")
	fmt.Fprintln(w)

	for _, table := range tables {
		writeTable(w, table)
	}
}

func parseTables(path string) ([]tableDef, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		tables   []tableDef
		current  *tableDef
		inCreate bool
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "-- ") {
			continue
		}

		if !inCreate {
			m := createTableRE.FindStringSubmatch(line)
			if m == nil {
				continue
			}
			inCreate = true
			current = &tableDef{name: m[1]}
			continue
		}

		if strings.HasPrefix(line, ")") && strings.HasSuffix(line, ";") {
			tables = append(tables, *current)
			current = nil
			inCreate = false
			continue
		}

		line = strings.TrimSuffix(line, ",")

		switch {
		case pkRE.MatchString(line):
			current.primaryKey = parseColumns(pkRE.FindStringSubmatch(line)[1])
		case indexRE.MatchString(line):
			m := indexRE.FindStringSubmatch(line)
			current.indexes = append(current.indexes, indexDef{
				name:    m[2],
				columns: parseColumns(m[3]),
				unique:  strings.TrimSpace(m[1]) != "",
			})
		case fulltextRE.MatchString(line):
			m := fulltextRE.FindStringSubmatch(line)
			current.indexes = append(current.indexes, indexDef{
				name:     m[1],
				columns:  parseColumns(m[2]),
				fulltext: true,
			})
		case columnRE.MatchString(line):
			m := columnRE.FindStringSubmatch(line)
			current.columns = append(current.columns, parseColumn(m[1], m[2]))
		default:
			return nil, fmt.Errorf("unhandled line in %s: %s", current.name, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

func parseColumn(name, raw string) columnDef {
	col := columnDef{
		name:          name,
		rawType:       firstField(raw),
		mappedType:    mapType(firstField(raw)),
		notNull:       strings.Contains(raw, "NOT NULL"),
		autoIncrement: strings.Contains(raw, "AUTO_INCREMENT"),
	}
	if m := regexp.MustCompile(`\bDEFAULT\s+(.+?)(?:\s+(?:NOT NULL|NULL|AUTO_INCREMENT)\b|$)`).FindStringSubmatch(raw); m != nil {
		col.defaultValue = strings.TrimSpace(m[1])
	}
	return col
}

func writeTable(w *bufio.Writer, table tableDef) {
	autoPK := singleAutoIncrementPK(table)

	fmt.Fprintf(w, "CREATE TABLE \"%s\" (\n", table.name)

	var defs []string
	for _, col := range table.columns {
		defs = append(defs, "  "+renderColumn(col, autoPK))
	}
	if len(table.primaryKey) > 0 && autoPK == "" {
		defs = append(defs, fmt.Sprintf("  PRIMARY KEY (%s)", quotedColumns(table.primaryKey)))
	}
	fmt.Fprintln(w, strings.Join(defs, ",\n"))
	fmt.Fprintln(w, ");")

	for _, idx := range table.indexes {
		if idx.fulltext {
			fmt.Fprintf(w, "-- Skipped FULLTEXT index \"%s\" on \"%s\" (%s)\n", idx.name, table.name, quotedColumns(idx.columns))
			continue
		}
		sqliteIndexName := fmt.Sprintf("%s__%s", table.name, idx.name)
		kind := "CREATE INDEX"
		if idx.unique {
			kind = "CREATE UNIQUE INDEX"
		}
		fmt.Fprintf(w, "%s \"%s\" ON \"%s\" (%s);\n", kind, sqliteIndexName, table.name, quotedColumns(idx.columns))
	}
	fmt.Fprintln(w)
}

func renderColumn(col columnDef, autoPK string) string {
	if col.name == autoPK {
		return fmt.Sprintf("\"%s\" INTEGER PRIMARY KEY AUTOINCREMENT", col.name)
	}

	parts := []string{
		fmt.Sprintf("\"%s\"", col.name),
		col.mappedType,
	}
	if col.notNull {
		parts = append(parts, "NOT NULL")
	}
	if col.defaultValue != "" {
		parts = append(parts, "DEFAULT", normalizeDefault(col))
	}
	return strings.Join(parts, " ")
}

func singleAutoIncrementPK(table tableDef) string {
	if len(table.primaryKey) != 1 {
		return ""
	}
	for _, col := range table.columns {
		if col.name == table.primaryKey[0] && col.autoIncrement {
			return col.name
		}
	}
	return ""
}

func parseColumns(raw string) []string {
	parts := strings.Split(raw, ",")
	cols := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		part = strings.Trim(part, "`")
		if part != "" {
			cols = append(cols, part)
		}
	}
	return cols
}

func quotedColumns(cols []string) string {
	quoted := make([]string, 0, len(cols))
	for _, col := range cols {
		quoted = append(quoted, fmt.Sprintf("\"%s\"", col))
	}
	return strings.Join(quoted, ", ")
}

func mapType(raw string) string {
	base := strings.ToLower(raw)
	switch {
	case base == "bool" || strings.HasPrefix(base, "tinyint"):
		return "INTEGER"
	case strings.HasPrefix(base, "int"), strings.HasPrefix(base, "smallint"), strings.HasPrefix(base, "bigint"):
		return "INTEGER"
	case strings.HasPrefix(base, "varchar"), strings.HasPrefix(base, "char"), strings.HasSuffix(base, "text"):
		return "TEXT"
	case strings.HasPrefix(base, "varbinary"), strings.HasPrefix(base, "binary"), strings.HasSuffix(base, "blob"):
		return "BLOB"
	default:
		return "TEXT"
	}
}

func normalizeDefault(col columnDef) string {
	if col.mappedType == "BLOB" && col.defaultValue == "\"\"" {
		return "X''"
	}
	return col.defaultValue
}

func firstField(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
