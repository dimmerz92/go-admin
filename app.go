package admin

import "net/http"

// App is the top-level application construct that bootstraps the underlying admin dashboard and admin public API.
type App struct {
	middleware []MiddlewareFunc
	mux        *http.ServeMux
}

type Config struct{}

// New returns a new default instance of the App with admin dashboard and public API bootstrapped.
func New() *App {
	return NewWithConfig(Config{})
}

// NewWithConfig returns a configured instance of the App with admin dashboard and public API bootstrapped.
func NewWithConfig(config Config) *App {
	_ = config
	return &App{
		mux: http.NewServeMux(),
	}
}

// Use registers global middleware.
// Note: Use should be called before any handlers are registered.
func (a *App) Use(middleware ...MiddlewareFunc) {
	a.middleware = append(a.middleware, middleware...)
}

// Handle registers the handler for the given pattern.
func (a *App) Handle(pattern string, handler http.Handler, middleware ...MiddlewareFunc) {
	a.mux.Handle(pattern, Chain(handler, append(a.middleware, middleware...)...))
}

// HandleFunc registers the handlerFunc for the given pattern.
func (a *App) HandleFunc(pattern string, handlerFunc http.HandlerFunc, middleware ...MiddlewareFunc) {
	a.Handle(pattern, handlerFunc, middleware...)
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
func (s *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
