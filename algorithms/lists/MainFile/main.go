package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) // 读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read {
		// 循环每个文件或者文件夹
		if fi.IsDir() {
			// 判断是否文件夹
			fullDir := path + "/" + fi.Name() // 构造新的路径
			files = append(files, fullDir)    // 追加路径
			files, _ = GetAll(fullDir, files) // 文件夹递归处理
		} else {
			fullDir := path + "/" + fi.Name()
			files = append(files, fullDir)
		}
	}
	return files, nil
}

func main() {
	path := "."
	files := make([]string, 0)        // 数组字符串
	files, _ = GetAll(path, files)    // 抓取所有文件
	for i := 0; i < len(files); i++ { // 打印路径
		fmt.Println(files[i])
	}
}
