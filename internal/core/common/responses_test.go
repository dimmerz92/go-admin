package common_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimmerz92/go-admin/internal/core/common"
)

func TestHTML(t *testing.T) {
	tests := []struct {
		name   string
		html   string
		status int
	}{
		{name: "status 200, no html", html: "", status: http.StatusOK},
		{name: "status 400, no html", html: "", status: http.StatusBadRequest},
		{name: "status 200, with html", html: "<p>Hello World</p>", status: http.StatusOK},
		{name: "status 400, with html", html: "<p>Hello World</p>", status: http.StatusBadRequest},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			if err := common.HTML(w, r, test.status, test.html); err != nil {
				t.Fatalf("failed to write HTML: %v", err)
			}

			if status := w.Result().StatusCode; status != test.status {
				t.Errorf("expected status %d, got %d", test.status, status)
			}

			if header := w.Header().Get("Content-Type"); header != "text/html" {
				t.Errorf("expected header text/html, got %s", header)
			}

			if body := w.Body.String(); body != test.html {
				t.Errorf("expected body %q, got %q", test.html, body)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     any
		expected string
		status   int
		err      bool
	}{
		{
			name:     "status 200, no data",
			data:     nil,
			expected: "null\n",
			status:   http.StatusOK,
			err:      false,
		},
		{
			name:     "status 400, no data",
			data:     nil,
			expected: "null\n",
			status:   http.StatusBadRequest,
			err:      false,
		},
		{
			name:     "status 200, with data",
			data:     map[string]string{"hello": "world"},
			expected: "{\"hello\":\"world\"}\n",
			status:   http.StatusOK,
			err:      false,
		},
		{
			name:     "status 400, with data",
			data:     map[string]string{"hello": "world"},
			expected: "{\"hello\":\"world\"}\n",
			status:   http.StatusBadRequest,
			err:      false,
		},
		{
			name:     "invalid data should error",
			data:     make(chan string),
			expected: "",
			status:   http.StatusOK,
			err:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			err := common.JSON(w, r, test.status, test.data)
			if test.err {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to write HTML: %v", err)
			}

			if status := w.Result().StatusCode; status != test.status {
				t.Errorf("expected status %d, got %d", test.status, status)
			}

			if header := w.Header().Get("Content-Type"); header != "application/json" {
				t.Errorf("expected header application/json, got %s", header)
			}

			if body := w.Body.String(); body != test.expected {
				t.Errorf("expected body %q, got %q", test.expected, body)
			}
		})
	}
}

func TestText(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		status int
	}{
		{name: "status 200, no text", text: "", status: http.StatusOK},
		{name: "status 400, no text", text: "", status: http.StatusBadRequest},
		{name: "status 200, with text", text: "Hello World", status: http.StatusOK},
		{name: "status 400, with text", text: "Hello World", status: http.StatusBadRequest},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			if err := common.Text(w, r, test.status, test.text); err != nil {
				t.Fatalf("failed to write HTML: %v", err)
			}

			if status := w.Result().StatusCode; status != test.status {
				t.Errorf("expected status %d, got %d", test.status, status)
			}

			if header := w.Header().Get("Content-Type"); header != "text/plain" {
				t.Errorf("expected header application/json, got %s", header)
			}

			if body := w.Body.String(); body != test.text {
				t.Errorf("expected body %q, got %q", test.text, body)
			}
		})
	}
}

func TestRedirect(t *testing.T) {
	tests := []struct {
		name           string
		status         int
		path           string
		htmx           bool
		expectedStatus int
		expectedHeader string
	}{
		{
			name:           "non htmx request",
			status:         http.StatusSeeOther,
			path:           "/redirect",
			htmx:           false,
			expectedStatus: http.StatusSeeOther,
			expectedHeader: "Location",
		},
		{
			name:           "htmx request",
			status:         http.StatusSeeOther,
			path:           "/redirect",
			htmx:           true,
			expectedStatus: http.StatusOK,
			expectedHeader: "Hx-Redirect",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			if test.htmx {
				r.Header.Set("Hx-Request", "true")
			}

			common.Redirect(w, r, test.status, test.path)

			if status := w.Result().StatusCode; status != test.expectedStatus {
				t.Errorf("expected status %d, got %d", test.expectedStatus, status)
			}

			if header := w.Header().Get(test.expectedHeader); header != test.path {
				t.Errorf("expected header %s to be %s, got %s", test.expectedHeader, test.path, header)
			}
		})
	}
}
