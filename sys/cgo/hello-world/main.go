package main

import "C"

/**
一个最简CGO程序该是什么样
要构造一个最简CGO程序，首先要忽视一些复杂的CGO特性，同时要展示CGO程序和纯程序的差别来！
*/

func main() {
	println("hello cgo")
}
