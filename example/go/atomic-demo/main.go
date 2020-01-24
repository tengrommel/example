package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	//wg.Add(100000)
	//for i:=0;i<100000;i++ {
	//	go add()
	//}
	//wg.Wait()
	//println(x)
	ok := atomic.CompareAndSwapInt64(&x, 0, 200)
	fmt.Println(ok, x)
}
