package web

import (
	"fmt"
	"net/http"
)

type SdkHttp interface {
	Route(pattern string, funcName func(*Context))
	Start(tag string)
	Shutdown()
}

type Server struct {
	name string
	context *Context
	node Filter
	route *Tree
}

func NewServer(name string, filter... FilterBuilder) *Server{
	routeMap := NewTree()
	c := NewContext()
	var base Filter = routeMap.ServeHTTP
	for _, f := range filter {
		base = f(base)
	}

	return &Server{
		name,
		c,
		base,
		routeMap,
	}
}

func (s *Server) Route(method, pattern string, funcName HandleMethod) {
	if root, ok := s.route.root[method]; !ok {
		fmt.Println("no support method")
		panic("no support method")
	}else {
		root.addNode(pattern, funcName)
	}
}

func (s *Server) Start(tag string) {
	http.ListenAndServe(tag, s)
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := &Context{writer, request}
	s.node(c)
}

func (s *Server) Shutdown() {
	fmt.Printf("%s关闭了", s.name)
}
