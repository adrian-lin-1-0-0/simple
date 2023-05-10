# Simple

Simple is a simple web framework written in Go


## Getting started


```bash
github.com/adrian-lin-1-0-0/simple
```

### A Simple Example


```go
package main

import (
	"github.com/adrian-lin-1-0-0/simple"
)

func main() {
	r := simple.Default()
	r.GET("/hello", func(c *simple.Context) {
		c.HTML("<h1>Hello</h1>")
	})

	v1 := r.Group("/v1")

	v1.GET("/hello", func(c *simple.Context) {
		a := []int{1, 2, 3}
		println(a[3])
		c.HTML("<h1>Hello v1</h1>")
	})

	v2 := r.Group("/v2")

	v2.GET("/hello", func(c *simple.Context) {
		c.HTML("<h1>Hello v2</h1>")
	})

	r.Run(":9999", func() {
		println("Server is running on port http://localhost:9999")
	})
}

```

## Class Diagram

```mermaid
---
title: Simple Web Framework 
---
classDiagram

    note for Engine"
        net/http

        func ListenAndServe(addr string, handler Handler) error {
            server := &Server{Addr: addr, Handler: handler}
            return server.ListenAndServe()
        }

        type Handler interface {
            ServeHTTP(ResponseWriter, *Request)
        }
    "

    class Engine{
        + ServeHTTP(http.ResponseWriter, *http.Request)
    }

    note for router "
        method - http method , e.g. : 
        GET,POST,PUT,DELETE
    "

    class router{
        - handler(c : *Context)
        - addRouter(method : String , fullPath : String ,handlers : HandlerFunc[])
    }

    class Context{
        + Writer : http.ResponseWriter
        + Req : *http.Request 
        + Params : HashMap<String,String>
        - handlers : HandlerFunc[]
    }

    note for node "
        Trie tree node.
        As shown in the diagram below (Trie Router).
    "

    class node{
        - fullPath : String
        - path : String
        - child : *node[]
        - wildChild : Boolean
        - handlers : HandlerFunc[]
    }

    note for group "
    All groups share the same Engine
    "

    class group{
        - middleware : HandlerFunc[]
        - parent : *group
        - engine : *Engine
    }

    class HandlerFunc{
        <<interface>> 
        - handler(c : *Context)
    }

    node "1" o-- "0..*" node
    group "1" o-- "0..*" group
    router "1" o-- "0..*" node
    Engine "1" o-- "1" router
    router ..> Context
    router ..> HandlerFunc
    group "1" o-- "0..*" HandlerFunc
    group "0..*" o-- "1" Engine
    HandlerFunc ..> Context
    Context "1" o-- "0..*" HandlerFunc
    
```

```mermaid
---
title: Trie Router
---
graph TB
    / --->|api| /api
    / --->|productpage| /productpage
    / --->|details| /details
    / --->|reviews| /reviews
    /reviews --->|:product_id| /reviews/:product_id
    /details --->|:product_id| /details/:product_id
    /api --->|v1| /v1  
```