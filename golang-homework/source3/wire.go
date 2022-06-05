// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"source3/internal/myapp/biz"
	"source3/internal/myapp/data"
	"source3/internal/myapp/service"
)

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) Servers{
	wire.Build(service.NewService, biz.NewBizobject, data.NewDB)
	return Servers  //返回值没有实际意义，只需符合函数签名即可
}