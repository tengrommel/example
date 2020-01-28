package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
	for {
		fmt.Println("周琳")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			return
		default:
		}
	}
}

/*
方法一： 使用全局变量
方法二： 使用chan
*/
func main() {
	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)
	exitChan <- true
	wg.Wait()
	// 如何通知子goroutine退出
}
