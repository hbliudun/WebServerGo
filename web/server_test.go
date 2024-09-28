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
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	h.addRouter(http.MethodGet, "/handle/bbb", func(ctx *Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	h.Get("/handle/aa", func(ctx *Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	h.Start(":8081")
}
