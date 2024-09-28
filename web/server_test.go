package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	h := NewHttpServer()
	//h := &HttpServer{}
	h.addRouter(http.MethodGet, "/user", func(ctx *Context) {
		fmt.Println("handle lam1")
		fmt.Println("handle lam2")
	})

	handler1 := func(ctx *Context) {
		fmt.Println("handler1")
	}
	handler2 := func(ctx *Context) {
		fmt.Println("handler 2")
	}

	h.addRouter(http.MethodGet, "/handle/bbb", func(ctx *Context) {
		handler2(ctx)
	})

	h.Get("/handle/aa", func(ctx *Context) {
		handler1(ctx)
	})

	h.Start(":8081")
}
