package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

type Cache struct {
	m map[string]string
	sync.Mutex
}

func (p *Cache) Lookup(key string) string {
	p.Lock()
	defer p.Unlock()
	return p.m[key]
}

// 具名函数
func Add(a, b int) int {
	return a + b
}

type Point struct {
	X, Y float64
}

// 匿名函数
var AddFunc = func(a, b int) int { return a + b }

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more对应[]int切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// 如果返回值命名了，可以通过名字来修改返回值，
// 也可以通过defer语句在return语句之后修改返回值
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

type File struct {
	fd int
}

func OpenFile(name string) (f *File, err error) {
	return nil, nil
}

func CloseFile(f *File) error {
	return nil
}

func ReadFile(f *File, offset int64, data []byte) int {
	return 0
}

func (f *File) Close() error {
	return nil
}

func (f *File) Read(int64, []byte) int {
	return 0
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	var a = []interface{}{123, "abc"}
	Print(a)
	Print(a...)
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
	for i := 0; i < 3; i++ {
		i := i
		defer func() { println(i) }()
	}
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
