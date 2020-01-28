package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
	for {
		fmt.Println("周琳")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 返回只读的chan
			return
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("OK")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 返回只读的chan
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
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go f(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	// 如何通知子goroutine退出
}
