package simple

import (
	"net/http"
)

type Engine struct {
	router *router
	*group
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.group = &group{engine: engine}
	return engine
}

// Run starts the HTTP server and listens for incoming requests.
// The provided callback function is called for each incoming request.
func (e *Engine) Run(addr string, callback func()) error {
	callback()
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	context.handlers = append(context.handlers, e.group.middleware...)
	e.router.handler(context)
}

func (e *Engine) AddRoute(method string, path string, handler HandlerFunc) {
	e.router.addRoute(method, path, []HandlerFunc{handler})
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodGet, path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodPost, path, handler)
}

func (e *Engine) DELETE(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodDelete, path, handler)
}

func (e *Engine) PATCH(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodPatch, path, handler)
}

func (e *Engine) PUT(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodPut, path, handler)
}

func (e *Engine) OPTIONS(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodOptions, path, handler)
}

func (e *Engine) HEAD(path string, handler HandlerFunc) {
	e.AddRoute(http.MethodHead, path, handler)
}
