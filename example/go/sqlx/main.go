package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	// 连接数据库
	dsn := "root:teng@tcp(127.0.0.1:3306)/imooc"
	// 连接数据库 db为数据库连接池
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		return fmt.Errorf("open %s failed, err: %v\n", dsn, err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Err: %v", err)
	}
	fmt.Println("连接数据库成功")
	db.SetMaxOpenConns(10)
	return nil
}

func main() {

}
