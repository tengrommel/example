package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	// 连接数据库
	dsn := "root:teng@tcp(127.0.0.1:3306)/imooc"
	// 连接数据库 db为数据库连接池
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return nil
}

type user struct {
	Id   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		return
	}
	sqlStr1 := `select id, name, age from user where id=3`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)
	var userList []user
	sqlStr2 := `select id, name, age from user`
	db.Select(&userList, sqlStr2)
	fmt.Printf("userList: %#v\n", userList)
}
