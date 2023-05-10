package simple

import "net/http"

type Simple struct {
	*group
}

func Default() *Simple {
	s := New()
	s.Use(Logger(), Recovery())
	return s
}

func New() *Simple {
	return &Simple{
		&group{
			prefix:     "/",
			router:     newRouter(),
			middleware: make(HandlersChain, 0),
		},
	}
}

func (s *Simple) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	context.handlers = append(context.handlers, s.group.middleware...)
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
