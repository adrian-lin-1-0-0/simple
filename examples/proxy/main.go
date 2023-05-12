package main

import (
	"log"

	"github.com/adrian-lin-1-0-0/simple"
)

func main() {
	proxy := simple.DefaultProxy()
	proxy.GET("/api/v1", "http://localhost:8888/v1")
	proxy.GET("/api/v2", "http://localhost:7777/v2")

	log.Fatal(
		proxy.Run(":9999", func() {
			println("Server is running on port http://localhost:9999")
		}),
	)
}
