package main

import (
	"fmt"
	"reflect"
)

// 反射的TypeOf()
func reflectType(x interface{}) {
	t := reflect.TypeOf(x) // 拿到x的动态类型信息
	fmt.Printf("type: %v\n", t)
	//fmt.Printf("%T\n", x) // 原理就是用的反射 代码补全也是用的反射
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	// 指针类型的t.Name是空
}

type cat struct {
	name string
}

type person struct {
	name string
	age  uint8
}

func main() {
	reflectType(100)
	reflectType(false)
	reflectType("沙河")
	reflectType([3]int{1, 2, 3})
	reflectType(map[string]int{"oo": 1})

	// 测试自定义的结构体类型
	var c1 = cat{
		name: "花花",
	}
	var p1 = person{
		name: "豪杰",
		age:  18,
	}
	reflectType(c1)
	reflectType(p1)

	var a int32 = 100
	var b = "豪杰"
	var f float32 = 12.34
	reflectType(a)
	reflectType(b)
	reflectType(f)
	reflectType(&a)
	reflectType(&b)
	reflectType(&f)
	var sliceInt = []int{1, 2, 3}
	reflectType(sliceInt)
	reflectType([2]int{})
	reflectType(map[string]int{})
}
