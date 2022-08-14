package web

import "net/http"

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext() *Context{
	return &Context{
		nil,nil,
	}
}