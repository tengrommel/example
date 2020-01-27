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

// 插入数据
func insert() {
	// 1、写语句
	sqlStr := `insert into user(name, age) values("图朝阳", 28)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:")
	}
	// 如果是插入数据的操作，能够拿到插入数量的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err: %v\n", err)
		return
	}
	fmt.Println(id)
}

// 更新操作
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err: %v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

func deleteRowDemo(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err: %v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name, age) values(?,?)`
	stmt, err := db.Prepare(sqlStr) // 把SQL语句先发给MySql先编译下
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	// 后续只需要拿到stmt去执行一些操作
	defer stmt.Close()
	var m = map[string]int{
		"刘启强": 30,
		"网相继": 32,
		"田硕":  40,
		"白对接": 40,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err: %v", err)
	}
	//fmt.Println("连接数据库成功！")
	//queryMore(0)
	//insert()
	//updateRow(8999, 2)
	prepareInsert()
}
