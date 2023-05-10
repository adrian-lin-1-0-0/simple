package simple

type router struct {
	methodTries map[string]*node
}

func newRouter() *router {
	return &router{
		methodTries: make(map[string]*node),
	}
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
