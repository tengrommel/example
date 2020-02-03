package main

import "fmt"

type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// 声明一个函数类型的变量，用于传参
type Option func(opts *Options)

// 初始化结构体
func InitOptions(opts ...Option) {
	options := &Options{}
	// 遍历opts，得到每一个函数
	for _, opt := range opts {
		// 调用函数，在函数里，给穿进去的对象赋值
		opt(options)
	}
	fmt.Printf("init options %#v\n", options)
}

// 定义具体给某个字段赋值的方法
func WithStrOption1(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}

// 定义具体给某个字段赋值的方法
func WithStrOption2(str string) Option {
	return func(opts *Options) {
		opts.strOption2 = str
	}
}

// 定义具体给某个字段赋值的方法
func WithStrOption3(str string) Option {
	return func(opts *Options) {
		opts.strOption3 = str
	}
}

// 定义具体给某个字段赋值的方法
func WithIntOption1(i int) Option {
	return func(opts *Options) {
		opts.intOption1 = i
	}
}

// 定义具体给某个字段赋值的方法
func WithIntOption2(i int) Option {
	return func(opts *Options) {
		opts.intOption2 = i
	}
}

// 定义具体给某个字段赋值的方法
func WithIntOption3(i int) Option {
	return func(opts *Options) {
		opts.intOption3 = i
	}
}

func main() {
	InitOptions(WithStrOption1("str1"), WithStrOption1("str2"))
}
