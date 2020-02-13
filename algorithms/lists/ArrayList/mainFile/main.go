package main

import (
	"awesomeProject/algorithms/lists/Queue"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "."
	files := make([]string, 0)
	myQueue := Queue.NewQueue()
	myQueue.EnQueue(path)
	for {
		path := myQueue.DeQueue() // 不断从队列中取出数据
		if path == nil {
			break
		}
		fmt.Println("get", path)
		read, _ := ioutil.ReadDir(path.(string))
		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path.(string) + "/" + fi.Name()
				fmt.Println("Dir", fullDir)
				myQueue.EnQueue(fullDir)
			} else {
				fullDir := path.(string) + "/" + fi.Name()
				files = append(files, fullDir)
				fmt.Println("File:", fullDir)
			}
		}
	}
}
