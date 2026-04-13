package forumsite

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mdhender/gobbs/internal/setupjson"
)

func TestSanitizeHTMLFragment(t *testing.T) {
	t.Parallel()

	got := string(sanitizeHTMLFragment(`<p>Hello <strong>world</strong><script>alert(1)</script><a href="javascript:alert(1)" title="x">bad</a><a href="https://example.com/path">good</a><img src="/img.png" onerror="alert(1)"></p>`))

	for _, unwanted := range []string{"<script", "javascript:alert(1)", "onerror"} {
		if strings.Contains(got, unwanted) {
			t.Fatalf("sanitized HTML contains %q: %s", unwanted, got)
		}
	}
	for _, wanted := range []string{"<p>Hello <strong>world</strong>", `<a title="x">bad</a>`, `<a href="https://example.com/path" rel="nofollow noreferrer noopener" target="_blank">good</a>`, `<img src="/img.png">`} {
		if !strings.Contains(got, wanted) {
			t.Fatalf("sanitized HTML missing %q: %s", wanted, got)
		}
	}
}

func TestRenderFormattedTextDefaultsToPlainText(t *testing.T) {
	t.Parallel()

	got := string(renderFormattedText(setupjson.TextFormatTaintedText, `<b>unsafe</b>`))
	if strings.Contains(got, "<b>unsafe</b>") {
		t.Fatalf("plain text rendering left raw HTML intact: %s", got)
	}
	if !strings.Contains(got, "&lt;b&gt;unsafe&lt;/b&gt;") {
		t.Fatalf("plain text rendering did not escape HTML: %s", got)
	}
}

func TestConfigTableNameIsUnquoted(t *testing.T) {
	t.Parallel()

	if got := configTableName("PBMnet_", "forums"); got != "PBMnet_forums" {
		t.Fatalf("configTableName() = %q, want %q", got, "PBMnet_forums")
	}
	if got := tableName("PBMnet_", "forums"); got != `"PBMnet_forums"` {
		t.Fatalf("tableName() = %q, want %q", got, `"PBMnet_forums"`)
	}
}

func TestContentDebugClass(t *testing.T) {
	t.Parallel()

	if got := contentDebugClass(setupjson.TextFormatRawHTML, true); got != " debug-raw-html" {
		t.Fatalf("contentDebugClass(raw-html, true) = %q", got)
	}
	if got := contentDebugClass(setupjson.TextFormatBBCodes, true); got != "" {
		t.Fatalf("contentDebugClass(bb-codes, true) = %q, want empty", got)
	}
	if got := contentDebugClass(setupjson.TextFormatRawHTML, false); got != "" {
		t.Fatalf("contentDebugClass(raw-html, false) = %q, want empty", got)
	}
}

func TestBuildWritesErrorPages(t *testing.T) {
	t.Parallel()

	renderer, err := New(Config{
		SQLitePath:   filepath.Join("..", "..", "mybb.sqlite3"),
		SetupPath:    filepath.Join("..", "..", "setup.json"),
		TemplatesDir: "templates",
		LiveTemplate: true,
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer renderer.Close()

	outDir := t.TempDir()
	if err := renderer.Build(outDir); err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	for _, tc := range []struct {
		name     string
		filename string
		contains []string
	}{
		{
			name:     "404 page",
			filename: "404.html",
			contains: []string{"404", "Page Not Found", "Return to the archive index"},
		},
		{
			name:     "500 page",
			filename: "500.html",
			contains: []string{"500", "Archive Error", "Return to the archive index"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(outDir, tc.filename))
			if err != nil {
				t.Fatalf("ReadFile(%q) error = %v", tc.filename, err)
			}
			html := string(data)
			for _, want := range tc.contains {
				if !strings.Contains(html, want) {
					t.Fatalf("%s missing %q", tc.filename, want)
				}
			}
		})
	}
}

func TestBuildCopiesUploadsDirectory(t *testing.T) {
	t.Parallel()

	uploadsDir := filepath.Join(t.TempDir(), "uploads")
	if err := os.MkdirAll(filepath.Join(uploadsDir, "avatars"), 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	want := "avatar-bytes"
	srcFile := filepath.Join(uploadsDir, "avatars", "avatar_1.png")
	if err := os.WriteFile(srcFile, []byte(want), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	renderer, err := New(Config{
		SQLitePath:   filepath.Join("..", "..", "mybb.sqlite3"),
		SetupPath:    filepath.Join("..", "..", "setup.json"),
		TemplatesDir: "templates",
		UploadsDir:   uploadsDir,
		LiveTemplate: true,
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer renderer.Close()

	outDir := t.TempDir()
	if err := renderer.Build(outDir); err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	got, err := os.ReadFile(filepath.Join(outDir, "uploads", "avatars", "avatar_1.png"))
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if string(got) != want {
		t.Fatalf("copied upload = %q, want %q", string(got), want)
	}
}

func TestBuildPrefixesAssetLinksWithBaseURL(t *testing.T) {
	t.Parallel()

	renderer, err := New(Config{
		SQLitePath:   filepath.Join("..", "..", "mybb.sqlite3"),
		SetupPath:    filepath.Join("..", "..", "setup.json"),
		TemplatesDir: "templates",
		BaseURL:      "/archive/",
		LiveTemplate: true,
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer renderer.Close()

	outDir := t.TempDir()
	if err := renderer.Build(outDir); err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	indexHTML, err := os.ReadFile(filepath.Join(outDir, "index.html"))
	if err != nil {
		t.Fatalf("ReadFile(index.html) error = %v", err)
	}
	if !strings.Contains(string(indexHTML), `href="/archive/assets/site.css"`) {
		t.Fatalf("index.html does not contain base-url-prefixed stylesheet link")
	}
}
