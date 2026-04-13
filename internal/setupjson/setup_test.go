package setupjson

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseTextFormats(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "setup.json")
	if err := os.WriteFile(path, []byte(`{
  "database": {
    "hostname": "localhost:3306",
    "database": "pbmnet",
    "table_prefix": "PBMnet_",
    "username": "user",
    "password": "pass"
  },
  "debug": {
    "highlightRawHTML": true
  },
  "textformats": {
    "PBMnet_forums": {
      "description": "raw-html"
    },
    "PBMnet_posts": {
      "message": "bb-codes"
    }
  }
}`), 0o644); err != nil {
		t.Fatalf("write setup.json: %v", err)
	}

	cfg, err := Parse(path)
	if err != nil {
		t.Fatalf("Parse(): %v", err)
	}

	if got := cfg.Database.TablePrefix; got != "PBMnet_" {
		t.Fatalf("table prefix = %q, want %q", got, "PBMnet_")
	}
	if !cfg.Debug.HighlightRawHTML {
		t.Fatal("expected debug.highlightRawHTML to be true")
	}
	if got := cfg.TextFormats.Format("PBMnet_forums", "description"); got != TextFormatRawHTML {
		t.Fatalf("description format = %q, want %q", got, TextFormatRawHTML)
	}
	if got := cfg.TextFormats.Format("PBMnet_posts", "message"); got != TextFormatBBCodes {
		t.Fatalf("message format = %q, want %q", got, TextFormatBBCodes)
	}
	if got := cfg.TextFormats.Format("PBMnet_threads", "subject"); got != TextFormatTaintedText {
		t.Fatalf("default format = %q, want %q", got, TextFormatTaintedText)
	}
}

func TestParseTextFormatsFailsFastOnUnknownFormat(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "setup.json")
	if err := os.WriteFile(path, []byte(`{
  "textformats": {
    "PBMnet_forums": {
      "description": "html"
    }
  }
}`), 0o644); err != nil {
		t.Fatalf("write setup.json: %v", err)
	}

	_, err := Parse(path)
	if err == nil {
		t.Fatal("Parse() error = nil, want validation failure")
	}
	if !strings.Contains(err.Error(), `textformats.PBMnet_forums.description: unknown format "html"`) {
		t.Fatalf("Parse() error = %q, want field-specific validation message", err)
	}
}
