package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	AddRoute(method string, path string, handlerFunc http.HandlerFunc)
}

type HttpServer struct {
	router
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: newRouter(),
	}
}

func (h *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.server(ctx)
}

func (h *HttpServer) server(ctx *Context) {
	// 干什么的呢？
	// 接下来就是查找路由，并且执行命中的业务逻辑
	n, ok := h.findRouter(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || n.handler == nil {
		ctx.Resp.WriteHeader(404)
		_, _ = ctx.Resp.Write([]byte("NOT FOUND"))
		return
	}
	n.handler(ctx)
}

func (h *HttpServer) Get(path string, handleFunc HandleFunc) {
	h.addRouter(http.MethodGet, path, handleFunc)
}

func (h *HttpServer) Post(path string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPost, path, handleFunc)
}

func (h *HttpServer) Option(path string, handleFunc HandleFunc) {
	h.addRouter(http.MethodOptions, path, handleFunc)
}

func (h *HttpServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h)
}
