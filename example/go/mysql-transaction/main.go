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

func transactionDemo() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err: %v\n", err)
		return
	}
	// 执行多个SQL操作
	sqlStr1 := `update user set age=age-2 where id=1`
	sqlStr2 := `update user set age=age+2 where id=3`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		return
	}
	// 上面两步SQL都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错啦， 要回滚！")
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("")
	}
	transactionDemo()
}
