package common

import (
	"encoding/json"
	"net/http"
)

func HTML(w http.ResponseWriter, r *http.Request, status int, html string) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)
	_, err := w.Write([]byte(html))
	return err
}

func JSON(w http.ResponseWriter, r *http.Request, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func Text(w http.ResponseWriter, r *http.Request, status int, text string) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	_, err := w.Write([]byte(text))
	return err
}

func IsHTMX(r *http.Request) bool {
	return r.Header.Get("Hx-Request") == "true"
}

func Redirect(w http.ResponseWriter, r *http.Request, status int, path string) {
	if IsHTMX(r) {
		w.Header().Set("Hx-Redirect", path)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, path, status)
}
