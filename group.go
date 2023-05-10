package simple

import "net/http"

type group struct {
	prefix     string
	middleware HandlersChain
	parent     *group
	router     *router
}

func (g *group) Group(prefix string, middleware ...HandlerFunc) *group {
	return &group{
		prefix:     g.prefix + prefix,
		middleware: append(g.middleware, middleware...),
		parent:     g,
		router:     g.router,
	}
}

func (g *group) Use(middleware ...HandlerFunc) *group {
	g.middleware = append(g.middleware, middleware...)
	return g
}

func (g *group) AddRoute(method, path string, handlers HandlersChain) {
	g.router.addRoute(method, g.prefix+path, handlers)
}

func (g *group) GET(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodGet, path, handlers)
}

func (g *group) POST(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodPost, path, handlers)
}

func (g *group) PUT(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodPut, path, handlers)
}

func (g *group) DELETE(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodDelete, path, handlers)
}

func (g *group) PATCH(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodPatch, path, handlers)
}

func (g *group) OPTIONS(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodOptions, path, handlers)
}

func (g *group) HEAD(path string, handlers ...HandlerFunc) {
	g.AddRoute(http.MethodHead, path, handlers)
}
