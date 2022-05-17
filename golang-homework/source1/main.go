package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type mysqlAction interface {
	writeSqlErr(sql string) error
}

func newconnect() *dbConfig {
	return &dbConfig{
		user: "root",
		password: "mysql_password",
		ip: "127.0.0.1",
		port: 3306,
		database: "test",
	}
}

type dbConfig struct {
	user string
	password string
	ip string
	port int
	database string
}

func (db *dbConfig)writeSqlErr(sqlString string) error {
	return errors.Wrapf(sql.ErrNoRows, "ip: %s, port: %d, database: %s, sql: %s", db.ip, db.port, db.database, sqlString)
}

func get() error{
	con := newconnect()
	if err := con.writeSqlErr("select * from test;"); err != nil{
		fmt.Println(err)
		return errors.WithMessage(err, "fail to get")
	}
	return nil
}

func main() {
	if err := get(); errors.Is(err, sql.ErrNoRows) {
	//if err := get(); errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("ERROR Types: %T\n", errors.Cause(err))
		fmt.Printf("ERROR Values: %v\n", err)
		fmt.Printf("ERROR tracing: %+v\n", err)
	}
}
