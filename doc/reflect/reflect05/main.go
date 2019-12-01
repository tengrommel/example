package main

import (
	"fmt"
	"reflect"
)

// 反射得到结构体的方法

type Student struct {
	name  string
	score int
}

func (s Student) Study(f int) string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是[]reflect.Value类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	var hj = Student{
		name:  "豪杰",
		score: 90,
	}
	printMethod(hj)
}
