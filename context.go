package simple

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
	query      url.Values
	index      int
	handlers   HandlersChain
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:     w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
		query:      req.URL.Query(),
		index:      -1,
		StatusCode: http.StatusOK,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(err string) {
	c.index = len(c.handlers)
	c.JSON(map[string]interface{}{"message": err})
}

func (c *Context) Query(key string) string {
	return c.query.Get(key)
}

func (c *Context) PostForm(key string) string {
	//need to refactor
	return c.Req.FormValue(key)
}

func (c *Context) Status(code int) *Context {
	c.StatusCode = code
	return c
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
	c.Writer.WriteHeader(c.StatusCode)
}

func (c *Context) String(format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) JSON(obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) Data(data []byte) {
	c.Writer.WriteHeader(c.StatusCode)
	_, err := c.Writer.Write(data)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) HTML(html string) {
	c.SetHeader("Content-Type", "text/html")
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
