package setupjson

import (
	"encoding/json"
	"fmt"
	"os"
)

type TextFormat string

const (
	TextFormatTaintedText TextFormat = "tainted-text"
	TextFormatRawHTML     TextFormat = "raw-html"
	TextFormatBBCodes     TextFormat = "bb-codes"
)

type Config struct {
	Database struct {
		Hostname    string `json:"hostname"`
		Database    string `json:"database"`
		TablePrefix string `json:"table_prefix"`
		Username    string `json:"username"`
		Password    string `json:"password"`
	} `json:"database"`
	Debug struct {
		HighlightRawHTML bool `json:"highlightRawHTML"`
	} `json:"debug"`
	TextFormats TextFormats `json:"textformats"`
}

type TextFormats map[string]map[string]TextFormat

func Parse(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if err := cfg.TextFormats.Validate(); err != nil {
		return Config{}, fmt.Errorf("parse %s: %w", path, err)
	}
	return cfg, nil
}

func (tf TextFormats) Format(table, column string) TextFormat {
	if cols, ok := tf[table]; ok {
		if format, ok := cols[column]; ok && format != "" {
			return format
		}
	}
	return TextFormatTaintedText
}

func (tf TextFormats) Validate() error {
	for table, cols := range tf {
		if table == "" {
			return fmt.Errorf("textformats: table name must not be empty")
		}
		for column, format := range cols {
			if column == "" {
				return fmt.Errorf("textformats.%s: column name must not be empty", table)
			}
			if !format.Valid() {
				return fmt.Errorf("textformats.%s.%s: unknown format %q", table, column, format)
			}
		}
	}
	return nil
}

func (f TextFormat) Valid() bool {
	switch f {
	case TextFormatTaintedText, TextFormatRawHTML, TextFormatBBCodes:
		return true
	default:
		return false
	}
}
