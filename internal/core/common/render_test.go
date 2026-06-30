package common_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimmerz92/go-admin/internal/core/common"
	"github.com/dimmerz92/go-admin/internal/mocks"
)

func TestRender(t *testing.T) {
	html := "<p>Hello World</p>"
	tests := []struct {
		name     string
		tpls     []common.Renderable
		status   int
		expected string
	}{
		{
			name:     "nil templates",
			tpls:     nil,
			status:   http.StatusOK,
			expected: "",
		},
		{
			name:     "empty templates",
			tpls:     []common.Renderable{},
			status:   http.StatusBadRequest,
			expected: "",
		},
		{
			name:     "one template",
			tpls:     []common.Renderable{mocks.Template(html)},
			status:   http.StatusTeapot,
			expected: html,
		},
		{
			name:     "many templates",
			tpls:     []common.Renderable{mocks.Template(html), mocks.Template(html), mocks.Template(html)},
			status:   http.StatusInternalServerError,
			expected: html + html + html,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			if err := common.Render(w, r, test.status, test.tpls...); err != nil {
				t.Fatalf("failed to render templates: %v", err)
			}

			if status := w.Result().StatusCode; status != test.status {
				t.Errorf("expected status %d, got %d", test.status, status)
			}

			if body := w.Body.String(); body != test.expected {
				t.Errorf("expected body %q, got %q", test.expected, body)
			}
		})
	}
}
