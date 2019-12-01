package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 是一个结构体， 它里面有一个计数器

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello 沙河！")
	if i == 8 {
		panic("报错了")
	}
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("hello main func.")
	wg.Wait()
	// 必须recover
	fmt.Println("main函数结束")
}
