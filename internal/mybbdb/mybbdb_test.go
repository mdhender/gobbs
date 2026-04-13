package mybbdb

import "testing"

func TestMysqlIdent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		{"users", "`users`"},
		{"my`table", "`my``table`"},
		{"", "``"},
		{"a``b", "`a````b`"},
	}
	for _, tt := range tests {
		if got := MysqlIdent(tt.input); got != tt.want {
			t.Errorf("MysqlIdent(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestSQLiteIdent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		{"users", `"users"`},
		{`my"table`, `"my""table"`},
		{"", `""`},
		{`a""b`, `"a""""b"`},
	}
	for _, tt := range tests {
		if got := SQLiteIdent(tt.input); got != tt.want {
			t.Errorf("SQLiteIdent(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestValidateConfig(t *testing.T) {
	t.Parallel()

	valid := Config{
		MysqlAddr:     "localhost:3306",
		MysqlDatabase: "mybb",
		MysqlUser:     "root",
	}
	if err := ValidateConfig(valid); err != nil {
		t.Fatalf("ValidateConfig(valid) = %v, want nil", err)
	}

	missing := []Config{
		{MysqlDatabase: "mybb", MysqlUser: "root"},
		{MysqlAddr: "localhost:3306", MysqlUser: "root"},
		{MysqlAddr: "localhost:3306", MysqlDatabase: "mybb"},
	}
	for _, cfg := range missing {
		if err := ValidateConfig(cfg); err == nil {
			t.Errorf("ValidateConfig(%+v) = nil, want error", cfg)
		}
	}
}
