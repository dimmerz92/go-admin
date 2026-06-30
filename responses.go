package admin

import (
	"net/http"

	"github.com/dimmerz92/go-admin/internal/core/common"
)

// HTML writes the status and html text to the response.
func HTML(w http.ResponseWriter, r *http.Request, status int, html string) error {
	return common.HTML(w, r, status, html)
}

// JSON writes the status and marshals the given data to the response.
func JSON(w http.ResponseWriter, r *http.Request, status int, data any) error {
	return common.JSON(w, r, status, data)
}

// Text writes the status and text to the response.
func Text(w http.ResponseWriter, r *http.Request, status int, text string) error {
	return common.Text(w, r, status, text)
}

// IsHTMX returns true if the request was issued by a HTMX element.
func IsHTMX(r *http.Request) bool {
	return common.IsHTMX(r)
}

// Redirect is a HTMX aware redirect.
// If the request originated from a HTMX element, the status is changed to 200 and a Hx-Redirect returned.
// https://github.com/bigskysoftware/htmx/issues/2052#issuecomment-1979805051
func Redirect(w http.ResponseWriter, r *http.Request, status int, path string) {
	common.Redirect(w, r, status, path)
}

// Renderable defines the contract for anything that can be rendered as HTML.
type Renderable = common.Renderable

// Render writes the status and any number of renderable HTML templates to the response.
// Recommended for Templ templates.
func Render(w http.ResponseWriter, r *http.Request, status int, tpls ...Renderable) error {
	return common.Render(w, r, status, tpls...)
}
