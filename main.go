package main

import (
	"WebServerGo/web"
	"fmt"
)

func main() {
	h := web.NewHttpServer()
	h.Get("/user", func(ctx *web.Context) {
		fmt.Println("handle lam1")
		fmt.Println("handle lam2")
	})

	handler1 := func(ctx *web.Context) {
		fmt.Println("handler1")
	}
	handler2 := func(ctx *web.Context) {
		fmt.Println("handler 2")
	}

	h.Get("/handle/bbb", func(ctx *web.Context) {
		handler2(ctx)
	})

	h.Get("/handle/aa", func(ctx *web.Context) {
		handler1(ctx)
	})

	h.Start(":8081")
}
