package admin_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dimmerz92/go-admin"
	"github.com/dimmerz92/go-admin/internal/mocks"
)

func TestApp(t *testing.T) {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			app := admin.New()
			app.Use(test.middleware...)
			app.HandleFunc("/", mocks.HandlerFunc("h", http.StatusOK))

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			app.ServeHTTP(w, r)

			if got := strings.Join(w.Header().Values("mw"), "|"); got != test.expected {
				t.Errorf("expected header %q, got %q", test.expected, got)
			}
		})
	}
}
