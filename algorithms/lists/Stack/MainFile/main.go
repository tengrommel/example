package main

import (
	"errors"
	"io/ioutil"
)

//
// --
// ----
func GetAllFileInPath(path string, files []string, level int) ([]string, error) {
	levelStr := ""
	if level == 1 {
		levelStr = "--"
	} else {
		for ; level >= 1; level-- {
			levelStr += "|--"
		}
	}
	readFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.New("文件夹不可读取")
	}
	for _, file := range readFiles {
		if file.IsDir() { // 判断是否文件夹
			fullDir := path + "/" + file.Name()
			files = append(files, levelStr+fullDir)
			files, _ = GetAllFileInPath(fullDir, files, level+1)
			// 递归处理
		} else {
			fullPath := path + "/" + file.Name()
			files = append(files, fullPath)
		}
	}
	return files, nil
}

func main() {

}
