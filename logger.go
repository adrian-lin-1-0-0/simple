package simple

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("%d %s %s in %v", c.StatusCode, c.Method, c.Req.RequestURI, time.Since(t))
	}
}
