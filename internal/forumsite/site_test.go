package forumsite

import (
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
