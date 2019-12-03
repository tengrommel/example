package db

import (
	myDB "awesomeProject/filestore-server/db/mysql"
	"fmt"
)

// UserSignUp:通过用户名及密码完成user表的注册操作
func UserSignUp(userName string, passWord string) bool {
	stmt, err := myDB.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`, `user_pwd`)values(?,?)")
	if err != nil {
		fmt.Println("Failed to insert, err:", err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(userName, passWord)
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

func UserSignIn(userName string, encPassword string) bool {
	stmt, err := myDB.DBConn().Prepare("select * from tbl_user where user_name=? limit  1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	rows, err := stmt.Query(userName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("username not found:" + userName)
		return false
	}
	pRows := myDB.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encPassword {
		return true
	}
	return false
}

func UpdateToken(userName string, token string) bool {
	stmt, err := myDB.DBConn().Prepare("replace into tbl_user_token (`user_name`, `user_token`) values (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(userName, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
