package data

import (
	"fmt"
	"source3/internal/myapp/biz"
)

type userMethod interface {
	Adduser(name string, age int32, opt ...string)
}

func (db *DBconfig) Adduser(name string, age int32, opt ...string) {
	_ = &biz.User{
		name,
		age,
		nil,nil,
	}
	fmt.Println("success to save user info")
}

type DBconfig struct {
	add string
	port int
	user string
	pass string
	db string
}

func NewDB() *DBconfig{
	return &DBconfig{}
}


