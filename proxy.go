package simple

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	*Simple
}

func DefaultProxy() *Proxy {
	return &Proxy{
		Simple: Default(),
	}
}

func NewProxy() *Proxy {
	return &Proxy{
		Simple: New(),
	}
}

func isValidUrl(u1 string) bool {

	_, err := url.ParseRequestURI(u1)
	if err != nil {
		return false
	}

	u, err := url.Parse(u1)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

func (p *Proxy) AddRoute(method, proxyPath, target string, handlers ...HandlerFunc) {

	if !isValidUrl(target) {
		panic(errors.New("invalid target url"))
	}

	proxyHandler := func(c *Context) {
		if len(c.Path) < len(proxyPath) {
			c.Status(http.StatusNotFound).
				String("%s: %s", http.StatusText(http.StatusNotFound), c.Path)
			return
		}

		targetPath := c.Path[len(proxyPath):]
		targetUrl, err := url.Parse(fmt.Sprintf("%s/%s", target, targetPath))
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(targetUrl)
		c.Req.URL.Path = "/"
		proxy.ServeHTTP(c.Writer, c.Req)
	}

	p.router.addRoute(method, proxyPath+"/*", append(handlers, proxyHandler))
}

func (p *Proxy) GET(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodGet, proxyPath, target, handlers...)
}

func (p *Proxy) POSR(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodPost, proxyPath, target, handlers...)
}

func (p *Proxy) PUT(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodPut, proxyPath, target, handlers...)
}

func (p *Proxy) DELETE(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodDelete, proxyPath, target, handlers...)
}

func (p *Proxy) PATCH(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodPatch, proxyPath, target, handlers...)
}

func (p *Proxy) Options(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodOptions, proxyPath, target, handlers...)
}

func (p *Proxy) Head(proxyPath, target string, handlers ...HandlerFunc) {
	p.AddRoute(http.MethodHead, proxyPath, target, handlers...)
}
