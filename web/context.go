package web

import "net/http"

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

//func AddRoute()
