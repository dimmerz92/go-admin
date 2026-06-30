package mocks_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimmerz92/go-admin/internal/mocks"
)

func TestHandler(t *testing.T) {
	tests := map[string]int{
		"ok":                    http.StatusOK,
		"accepted":              http.StatusAccepted,
		"bad request":           http.StatusBadRequest,
		"internal server error": http.StatusInternalServerError,
	}

	for text, status := range tests {
		t.Run(text, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			mocks.Handler(text, status).ServeHTTP(w, r)

			if got := w.Result().StatusCode; got != status {
				t.Errorf("expected status %d, got %d", status, got)
			}

			if got := w.Body.String(); got != text {
				t.Errorf("expected body %q, got %q", text, got)
			}
		})
	}
}

func TestHandlerFunc(t *testing.T) {
	tests := map[string]int{
		"ok":                    http.StatusOK,
		"accepted":              http.StatusAccepted,
		"bad request":           http.StatusBadRequest,
		"internal server error": http.StatusInternalServerError,
	}

	for text, status := range tests {
		t.Run(text, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			mocks.HandlerFunc(text, status).ServeHTTP(w, r)

			if got := w.Result().StatusCode; got != status {
				t.Errorf("expected status %d, got %d", status, got)
			}

			if got := w.Body.String(); got != text {
				t.Errorf("expected body %q, got %q", text, got)
			}
		})
	}
}

func TestMiddlewareFunc(t *testing.T) {
	m := mocks.MiddlewareFunc("mw", "mw")
	h := mocks.Handler("h", http.StatusTeapot)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	m(h).ServeHTTP(w, r)

	if got := w.Result().StatusCode; got != http.StatusTeapot {
		t.Errorf("expected status %d, got %d", http.StatusTeapot, got)
	}

	if got := w.Header().Get("mw"); got != "mw" {
		t.Errorf("expected header: %q, got %q", "mw", got)
	}

	if got := w.Body.String(); got != "h" {
		t.Errorf("expected body %q, got %q", "h", got)
	}
}
