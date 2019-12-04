package db

import (
	myDB "awesomeProject/filestore-server/db/mysql"
	"fmt"
	"time"
)

// UserFile: 用户文件表结构体
type UserFile struct {
	UserName    string
	FileHash    string
	FileName    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

func OnUserFileUploadFinished(userName, fileHash, fileName string, fileSize int64) bool {
	stmt, err := myDB.DBConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`, `file_sha1`, `file_name`," +
			"`file_size`, `upload_at`)values(?,?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(userName, fileHash, fileName, fileSize, time.Now())
	if err != nil {
		return false
	}
	return true
}

// QueryUserFileMetas批量获取用户文件信息
func QueryUserFileMetas(userName string, limit int) ([]UserFile, error) {
	stmt, err := myDB.DBConn().Prepare(
		"select file_sha1, file_name, file_size, upload_at, last_update from tbl_user_file" +
			"where user_name=? limit=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userName, limit)
	if err != nil {
		return nil, err
	}
	var userFiles []UserFile
	for rows.Next() {
		uFile := UserFile{}
		rows.Scan(&uFile.FileHash, &uFile.FileName, &uFile.FileSize,
			&uFile.UploadAt, &uFile.LastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		userFiles = append(userFiles, uFile)
	}
	return userFiles, nil
}
