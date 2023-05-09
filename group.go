package simple

import "net/http"

type group struct {
	prefix     string
	middleware []HandlerFunc
	parent     *group
	engine     *Engine
}

func (g *group) Group(prefix string) *group {
	engine := g.engine
	newGroup := &group{
		prefix: prefix,
		parent: g,
		engine: engine,
	}

	return newGroup
}

func (g *group) Use(middleware ...HandlerFunc) *group {
	g.middleware = append(g.middleware, middleware...)
	return g
}

func (g *group) addRoute(method string, path string, handler HandlerFunc) {
	fullPath := g.prefix + path
	g.engine.router.addRoute(method, fullPath, append(g.middleware, handler))
}

func (g *group) GET(path string, handler HandlerFunc) {
	g.addRoute(http.MethodGet, path, handler)
}

func (g *group) POST(path string, handler HandlerFunc) {
	g.addRoute(http.MethodPost, path, handler)
}

func (g *group) DELETE(path string, handler HandlerFunc) {
	g.addRoute(http.MethodDelete, path, handler)
}

func (g *group) PATCH(path string, handler HandlerFunc) {
	g.addRoute(http.MethodPatch, path, handler)
}

func (g *group) PUT(path string, handler HandlerFunc) {
	g.addRoute(http.MethodPut, path, handler)
}

func (g *group) OPTIONS(path string, handler HandlerFunc) {
	g.addRoute(http.MethodOptions, path, handler)
}

func (g *group) HEAD(path string, handler HandlerFunc) {
	g.addRoute(http.MethodHead, path, handler)
}
