# Simple

Simple is a simple web framework written in Go



## Class Diagram

```mermaid
---
title: Simple Web Framework 
---
classDiagram

    note for Simple"

        - Simple and all groups share the same router.

        - Simple need to implement the ServeHTTP method of the Handler interface.

        net/http

            func ListenAndServe(addr string, handler Handler) error {
                server := &Server{Addr: addr, Handler: handler}
                return server.ListenAndServe()
            }

            type Handler interface {
                ServeHTTP(ResponseWriter, *Request)
            }

    "

    class Simple{
        - handler(c : *Context)
        + ServeHTTP(http.ResponseWriter, *http.Request)
    }

    note for group "
    All groups share the same router
    "

    class group{
        - middleware : HandlerFunc[]
        - parent : *group
        - router : *router
    }
    
    class router{
        - addRouter(method : String , fullPath : String ,handlers : HandlerFunc[])
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

    class Context{
        + Writer : http.ResponseWriter
        + Req : *http.Request 
        + Params : HashMap<String,String>
        - handlers : HandlerFunc[]
    }

    class HandlerFunc{
        <<interface>> 
        - handler(c : *Context)
    }

    Simple "1" o-- "1" group 
    Simple "1" o-- "1" router 
    group "1" o-- "0..*" group
    node "1" o-- "0..*" node
    group "0..1" o-- "1*" router
    router "1" o-- "0..*" node
    node ..> Context
    Simple ..> Context
    Context o--> HandlerFunc
    HandlerFunc ..> Context
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


## Getting started

### Installation

```sh
go get github.com/adrian-lin-1-0-0/simple
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
		//recovery
		println(a[3]) // panic: runtime error: index out of range [3] with length 3
		c.HTML("<h1>Hello v1</h1>")
	})

	v2 := r.Group("/v2")

	v2.GET("/hello", func(c *simple.Context) {
		c.HTML("<h1>Hello v2</h1>")
	})

	r.Run(":8888", func() {
		println("Server is running on port http://localhost:8888")
	})
}
```