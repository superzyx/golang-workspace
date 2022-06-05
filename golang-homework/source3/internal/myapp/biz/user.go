package biz

import "source3/internal/myapp/data"


type User struct{
	Name string `json:"name"`
	Age int32 `json:"age"`
	Description string `json:"description"`
	Address string `json:"address"`
}

func NewBizobject(db *data.DBconfig) *User{
	return &User{
	}
}