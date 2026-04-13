package main

import "testing"

func TestParseColumn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		raw  string
		want columnDef
	}{
		// NOT NULL detection
		{
			name: "notNull",
			raw:  "int(11) NOT NULL",
			want: columnDef{
				name:       "notNull",
				rawType:    "int(11)",
				mappedType: "INTEGER",
				notNull:    true,
			},
		},
		// NULL (without NOT) should not set notNull
		{
			name: "nullable",
			raw:  "varchar(255) NULL",
			want: columnDef{
				name:       "nullable",
				rawType:    "varchar(255)",
				mappedType: "TEXT",
				notNull:    false,
			},
		},
		// AUTO_INCREMENT detection
		{
			name: "autoInc",
			raw:  "int(11) NOT NULL AUTO_INCREMENT",
			want: columnDef{
				name:          "autoInc",
				rawType:       "int(11)",
				mappedType:    "INTEGER",
				notNull:       true,
				autoIncrement: true,
			},
		},
		// DEFAULT with a numeric value
		{
			name: "defNum",
			raw:  "int(11) NOT NULL DEFAULT 0",
			want: columnDef{
				name:         "defNum",
				rawType:      "int(11)",
				mappedType:   "INTEGER",
				notNull:      true,
				defaultValue: "0",
			},
		},
		// DEFAULT with a quoted string
		{
			name: "defStr",
			raw:  "varchar(100) NOT NULL DEFAULT ''",
			want: columnDef{
				name:         "defStr",
				rawType:      "varchar(100)",
				mappedType:   "TEXT",
				notNull:      true,
				defaultValue: "''",
			},
		},
		// DEFAULT followed by NOT NULL (order reversed)
		{
			name: "defBeforeNotNull",
			raw:  "varchar(50) DEFAULT 'yes' NOT NULL",
			want: columnDef{
				name:         "defBeforeNotNull",
				rawType:      "varchar(50)",
				mappedType:   "TEXT",
				notNull:      true,
				defaultValue: "'yes'",
			},
		},
		// No DEFAULT, no NOT NULL, no AUTO_INCREMENT
		{
			name: "bare",
			raw:  "text",
			want: columnDef{
				name:       "bare",
				rawType:    "text",
				mappedType: "TEXT",
			},
		},
		// DEFAULT at end of line with no trailing qualifier
		{
			name: "defAtEnd",
			raw:  "tinyint(1) DEFAULT 1",
			want: columnDef{
				name:         "defAtEnd",
				rawType:      "tinyint(1)",
				mappedType:   "INTEGER",
				defaultValue: "1",
			},
		},
	}
	for _, tt := range tests {
		got := parseColumn(tt.name, tt.raw)
		if got != tt.want {
			t.Errorf("parseColumn(%q, %q)\n  got  %+v\n  want %+v", tt.name, tt.raw, got, tt.want)
		}
	}
}

func TestNormalizeDefault(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		col  columnDef
		want string
	}{
		// BLOB with "" default → X''
		{
			name: "blobEmptyQuoted",
			col:  columnDef{mappedType: "BLOB", defaultValue: `""`},
			want: "X''",
		},
		// BLOB with a non-"" default passes through
		{
			name: "blobOtherDefault",
			col:  columnDef{mappedType: "BLOB", defaultValue: "'data'"},
			want: "'data'",
		},
		// Non-BLOB type with "" default passes through
		{
			name: "textEmptyQuoted",
			col:  columnDef{mappedType: "TEXT", defaultValue: `""`},
			want: `""`,
		},
		// INTEGER default passes through
		{
			name: "integerDefault",
			col:  columnDef{mappedType: "INTEGER", defaultValue: "0"},
			want: "0",
		},
		// Empty default passes through
		{
			name: "emptyDefault",
			col:  columnDef{mappedType: "TEXT", defaultValue: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		if got := normalizeDefault(tt.col); got != tt.want {
			t.Errorf("normalizeDefault(%s) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

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
