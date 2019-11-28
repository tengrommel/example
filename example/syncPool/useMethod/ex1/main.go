package main

import (
	"fmt"
	"sync"
)

func main() {
	// 如果我们不指定New函数的话，会返回nil
	p := &sync.Pool{New: func() interface{} {
		return 0
	}}
	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)
	c := p.Get().(int)
	fmt.Println(c)
}
