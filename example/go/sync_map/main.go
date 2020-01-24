package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go 内置的map不是并发安全的
var m = make(map[string]int)

var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n) // 必须使用sync.Map内置的store方法设置键值
			value, _ := m2.Load(key)
			fmt.Printf("k:=%v, v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
