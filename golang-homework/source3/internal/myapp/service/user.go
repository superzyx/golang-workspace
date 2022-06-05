package service

import (
	"context"
	"source3/internal/myapp/data"
	"source3/internal/myapp/biz"
)

type Servers struct {
	name string
	ctx context.Context
	cancel context.CancelFunc
	db *data.DBconfig
	user *biz.User
}

func NewServer(name string) *Servers{
	ctx := context.Background()
	ctx, cencel := context.WithCancel(ctx)
	return &Servers{
		name,
		ctx,
		cencel,
		nil, nil,
	}
}


func (s *Servers)UserSignUp(name string, age int32) {
	s.db.Adduser(name, age)
}

func NewService(db *data.DBconfig, user *biz.User) *Servers{
	s := NewServer("user")
	s.db = db
	s.user = user
	return s
}
