package main

import (
	"awesomeProject/algorithms/lists/Stack/StackArray"
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
	files := make([]string, 0)
	myStack := StackArray.NewStack()
	myStack.Push(path)
	for !myStack.IsEmpty() {
		getPath := myStack.Pop().(string)
		files = append(files, getPath)     // 加入列表
		read, _ := ioutil.ReadDir(getPath) // 读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := getPath + "/" + fi.Name()
				files = append(files, fullDir)
				myStack.Push(fullDir)
			} else {
				fullDir := getPath + "/" + fi.Name()
				files = append(files, fullDir)
			}
		}
	}

	for i := 0; i < len(files); i++ { // 打印
		fmt.Println(files[i])
	}

}
