package main

import "testing"

func TestMapType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		// bool / tinyint → INTEGER
		{"bool", "INTEGER"},
		{"BOOL", "INTEGER"},
		{"tinyint", "INTEGER"},
		{"tinyint(1)", "INTEGER"},
		{"TINYINT(4)", "INTEGER"},

		// int variants → INTEGER
		{"int", "INTEGER"},
		{"int(11)", "INTEGER"},
		{"INT(10)", "INTEGER"},
		{"smallint", "INTEGER"},
		{"smallint(6)", "INTEGER"},
		{"bigint", "INTEGER"},
		{"bigint(20)", "INTEGER"},

		// text types → TEXT
		{"varchar(255)", "TEXT"},
		{"VARCHAR(100)", "TEXT"},
		{"char(1)", "TEXT"},
		{"CHAR(32)", "TEXT"},
		{"text", "TEXT"},
		{"TEXT", "TEXT"},
		{"mediumtext", "TEXT"},
		{"longtext", "TEXT"},
		{"tinytext", "TEXT"},

		// blob / binary types → BLOB
		{"blob", "BLOB"},
		{"BLOB", "BLOB"},
		{"mediumblob", "BLOB"},
		{"longblob", "BLOB"},
		{"tinyblob", "BLOB"},
		{"varbinary(255)", "BLOB"},
		{"binary(16)", "BLOB"},

		// unknown falls back to TEXT
		{"decimal(10,2)", "TEXT"},
		{"date", "TEXT"},
		{"timestamp", "TEXT"},
		{"float", "TEXT"},
		{"double", "TEXT"},
		{"enum('a','b')", "TEXT"},
	}
	for _, tt := range tests {
		if got := mapType(tt.input); got != tt.want {
			t.Errorf("mapType(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
