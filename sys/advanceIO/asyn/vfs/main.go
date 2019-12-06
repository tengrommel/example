package main

import (
	"fmt"
	"time"
)

func worker(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cannel:
			// 退出
		}
	}
}

func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
	cannel := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(cannel)
	}
	time.Sleep(time.Second)
	close(cannel)
}
