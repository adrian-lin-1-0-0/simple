package simple

import (
	"log"
	"net/http"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%s\n", err)
				c.Status(http.StatusInternalServerError).
					Fail(http.StatusText(http.StatusInternalServerError))
			}
		}()

		c.Next()
	}
}
