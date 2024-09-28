package main

import (
	"WebServerGo/web"
	"fmt"
)

func main() {
	h := web.NewHttpServer()
	h.Get("/user", func(ctx *web.Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	h.Get("/handle/bbb", func(ctx *web.Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	h.Get("/handle/aa", func(ctx *web.Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})
	//h.Get("/temp/ccc", func(ctx *web.Context) {
	//	ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	//})
	h.Get("/*/abc", func(ctx *web.Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})
	h.Get("/temp/*", func(ctx *web.Context) {
		ctx.Resp.Write([]byte(fmt.Sprintf("hello, %s", ctx.Req.URL.Path)))
	})

	h.Start(":8081")
}
