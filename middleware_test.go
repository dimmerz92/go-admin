package admin_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dimmerz92/go-admin"
	"github.com/dimmerz92/go-admin/internal/mocks"
)

func TestChain(t *testing.T) {
	tests := []struct {
		name       string
		middleware []admin.MiddlewareFunc
		expected   string
	}{
		{name: "nil middleware", middleware: nil, expected: ""},
		{name: "empty middleware", middleware: []admin.MiddlewareFunc{}, expected: ""},
		{
			name:       "one middleware",
			middleware: []admin.MiddlewareFunc{mocks.MiddlewareFunc("mw", "m1")},
			expected:   "m1",
		},
		{
			name:       "three middleware",
			middleware: []admin.MiddlewareFunc{mocks.MiddlewareFunc("mw", "m1"), mocks.MiddlewareFunc("mw", "m2"), mocks.MiddlewareFunc("mw", "m3")},
			expected:   "m1|m2|m3",
		},
		{
			name:       "inverse three middleware",
			middleware: []admin.MiddlewareFunc{mocks.MiddlewareFunc("mw", "m3"), mocks.MiddlewareFunc("mw", "m2"), mocks.MiddlewareFunc("mw", "m1")},
			expected:   "m3|m2|m1",
		},
	}

	handler := mocks.Handler("h", http.StatusOK)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			admin.Chain(handler, test.middleware...).ServeHTTP(w, r)

			if got := w.Result().StatusCode; got != http.StatusOK {
				t.Errorf("expected status %d, got %d", http.StatusOK, got)
			}

			if got := strings.Join(w.Header().Values("mw"), "|"); got != test.expected {
				t.Errorf("expected header %q, got %q", test.expected, got)
			}

			if got := w.Body.String(); got != "h" {
				t.Errorf("expected body %q, got %q", "h", got)
			}
		})
	}
}
