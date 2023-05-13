package simple

import (
	"net/http"
	"strings"
)

type Simple struct {
	*group
	groups []*group
}

func Default() *Simple {
	s := New()
	s.Use(Logger(), Recovery())
	return s
}

func New() *Simple {

	s := &Simple{}
	g := &group{
		prefix:     "",
		router:     newRouter(),
		simple:     s,
		middleware: make(HandlersChain, 0),
	}
	s.group = g
	s.groups = []*group{g}
	return s
}

func (s *Simple) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc

	for _, group := range s.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleware...)
		}
	}
	context := newContext(w, r)
	context.handlers = middlewares
	s.handler(context)
}

func (s *Simple) Run(addr string, callback func()) error {
	callback()
	return http.ListenAndServe(addr, s)
}

func (s *Simple) handler(c *Context) {
	node := s.router.getRoute(c.Method, c.Path)
	if node == nil {
		c.handlers = append(c.handlers, func(c *Context) {
			c.Status(http.StatusNotFound).
				String("%s: %s", http.StatusText(http.StatusNotFound), c.Path)

		})
		goto Next
	}
	c.Params = node.parseParams(c.Path)
	c.handlers = append(c.handlers, node.handlers...)

Next:
	c.Next()
}
