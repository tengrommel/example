package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
	age  int
}

// 定义一个全局的变量
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

// 查询
func query() {
	sqlStr := `select id, name, age from user where id=?;`
	// 执行
	row := db.QueryRow(sqlStr, 2) // 从连接池里拿一个连接出来去数据库查询单条记录
	// 拿到结果
	var u user
	row.Scan(&u.id, &u.name, &u.age)
	fmt.Printf("u1: %#v\n", u)
}

// 查询多条
func queryMore(n int) {
	sqlStr := `select id, name, age from user where id > ?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("exec %s query failed, err: %v\n", sqlStr, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("u:%#v\n", u)
		}
		fmt.Printf("u:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err: %v", err)
	}
	fmt.Println("连接数据库成功！")
	queryMore(0)
}
