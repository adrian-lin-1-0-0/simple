package simple

import "strings"

type node struct {
	fullPath  string
	path      string
	child     []*node
	wildChild bool
	handlers  HandlersChain
}

type nodeOptions struct {
	fullPath  string
	path      string
	wildChild bool
	handlers  HandlersChain
}

func newNode(opts *nodeOptions) *node {
	if opts == nil {
		return &node{
			child: make([]*node, 0),
		}
	}

	return &node{
		child:     make([]*node, 0),
		fullPath:  opts.fullPath,
		path:      opts.path,
		wildChild: opts.wildChild,
		handlers:  opts.handlers,
	}
}

func (n *node) addRoute(fullPath string, handlers HandlersChain) {
	parts := fullPathToPaths(fullPath)
	n.insert(fullPath, parts, handlers)
}

func (n *node) getRoute(fullPath string) *node {
	parts := fullPathToPaths(fullPath)
	return n.search(parts)
}

func (n *node) findOneChild(path string) *node {
	for _, child := range n.child {
		if child.path == path || child.wildChild {
			return child
		}
	}
	return nil
}

func (n *node) insert(pattern string, parts []string, handlers HandlersChain) {
	curr := n
	for _, part := range parts {
		child := curr.findOneChild(part)
		if child == nil {
			child = newNode(&nodeOptions{
				path:      part,
				wildChild: part[0] == ':' || part[0] == '*',
			})
			curr.child = append(curr.child, child)
		}
		curr = child
	}
	curr.fullPath = pattern
	curr.handlers = handlers
}

func (n *node) IsWilcard() bool {
	return n.path[0] == '*'
}

func (n *node) search(parts []string) *node {
	child := n
	for _, part := range parts {
		child = child.findOneChild(part)
		if child == nil {
			return nil
		}
		if child.IsWilcard() {
			return child
		}
	}

	return child
}

func (n *node) parseParams(fullPath string) map[string]string {
	params := make(map[string]string)
	reqPaths := fullPathToPaths(fullPath)
	paths := fullPathToPaths(n.fullPath)
	for idx, path := range paths {

		if reqPaths[idx] != path && path[0] != ':' {
			break
		}
		params[path[1:]] = reqPaths[idx]
	}
	return params
}

func fullPathToPaths(fullPath string) []string {
	paths := []string{}
	for _, path := range strings.Split(fullPath, "/") {
		if path != "" {
			paths = append(paths, path)
		}
	}
	return paths
}
