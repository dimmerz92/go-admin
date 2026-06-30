package mocks_test

import (
	"bytes"
	"testing"

	"github.com/dimmerz92/go-admin/internal/mocks"
)

func TestTemplate(t *testing.T) {
	html := "<p>hello world</p>"
	tpl := mocks.Template(html)

	var buf bytes.Buffer
	if err := tpl.Render(t.Context(), &buf); err != nil {
		t.Fatalf("failed to render: %v", err)
	}

	if got := buf.String(); got != html {
		t.Fatalf("expected %q, got %q", html, got)
	}
}
