package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

//
// ---
// ------
func GetAll(path string, files []string, level int) ([]string, error) {
	fmt.Println("level", level)
	levelStr := ""
	if level == 1 {
		levelStr = "+"
	} else {
		for ; level > 1; level-- {
			levelStr += "|---"
		}
		levelStr += "+"
	}
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
			newLevel := level + 1
			files, _ = GetAll(fullDir, files, newLevel) // 文件夹递归处理
		} else {
			fullDir := path + "/" + fi.Name()
			files = append(files, levelStr+fullDir)
		}
	}
	return files, nil
}

func main() {
	// 深度

}
