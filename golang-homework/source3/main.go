package main

import (
	"context"
	"source3/internal/myapp/service"
)

type registerMethod interface {
	add(s ...*service.Servers)
	delete()
	update()
	get()
}

type ServerMethod interface {
	Start()
	Stop()
}

func (r *register) add(s ...*service.Servers) {
	r.server = append(r.server, s...)
}

func (r *register) delete() {
	panic("implement me")
}

func (r *register) update() {
	panic("implement me")
}

func (r *register) get() {
	panic("implement me")
}

func (s *service.Servers) Start() {
	panic("implement me")
}

func (s *service.Servers) Stop() {
	s.cancel()
}

type register struct {
	// 注册中心信息
	api string
	user string
	password string
	token string
	server []*service.Servers
}

func Newregister() *register{
	return &register{
		api: "https://xxxx/api/v1",
		server: make([]*service.Servers, 3),
	}
}

//type Servers struct {
//	name string
//	ctx context.Context
//	cancel context.CancelFunc
//}
//
//func NewServer(name string) *Servers{
//	ctx := context.Background()
//	ctx, cencel := context.WithCancel(ctx)
//	return &Servers{
//		name,
//		ctx,
//		cencel,
//	}
//}

func main() {
	rgt := Newregister()
	UserServer := service.NewServer("user")
	orderServer := service.NewServer("order")
	rgt.add(UserServer, orderServer)
}
