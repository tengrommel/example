package main

import (
	"awesomeProject/algorithms/lists/Stack/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归文件
func GetAllFileInPath(path string, files []string) ([]string, error) {
	readFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.New("文件夹不可读取")
	}
	for _, file := range readFiles {
		if file.IsDir() { // 判断是否文件夹
			fullDir := path + "/" + file.Name()
			files = append(files, fullDir)
			files, _ = GetAllFileInPath(fullDir, files) // 递归处理
		} else {
			fullPath := path + "/" + file.Name()
			files = append(files, fullPath)
		}
	}
	return files, nil
}

func main() {
	path := "/Users/tengzhou/git/awesomeProject/go-rest-app"
	files := make([]string, 0)
	myStack := StackArray.NewStack()
	myStack.Push(path)
	for !myStack.IsEmpty() {
		path := myStack.Pop().(string)
		files = append(files, path)
		read, _ := ioutil.ReadDir(path)
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path + "/" + fi.Name()
				files = append(files, fullDir)
				myStack.Push(fullDir)
			} else {
				fullDir := path + "/" + fi.Name()
				files = append(files, fullDir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
