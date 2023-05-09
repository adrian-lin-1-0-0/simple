package simple

import (
	"net/http"
)

type router struct {
	methodTries map[string]*node
}

func newRouter() *router {
	return &router{
		methodTries: make(map[string]*node),
	}
}

func (r *router) handler(c *Context) {

	node := r.getRoute(c.Method, c.Path)
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

func (r *router) addRoute(method, fullPath string, handlersChain HandlersChain) {
	_, ok := r.methodTries[method]
	if !ok {
		r.methodTries[method] = newNode(nil)
	}

	r.methodTries[method].addRoute(fullPath, handlersChain)
}

func (r *router) getRoute(method, fullPath string) *node {
	_, ok := r.methodTries[method]
	if !ok {
		return nil
	}

	return r.methodTries[method].getRoute(fullPath)
}
