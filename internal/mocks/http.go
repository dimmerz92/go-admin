package mocks

import "net/http"

// Handler writes the values and status to the response writer.
func Handler(value string, status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(value))
	})
}

// HandlerFunc writes the values and status to the response writer.
func HandlerFunc(value string, status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(value))
	}
}

// MiddlewareFunc returns a middleware that writes the header to the response and calls the next handler.
func MiddlewareFunc(header, value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(header, value)
			next.ServeHTTP(w, r)
		})
	}
}
