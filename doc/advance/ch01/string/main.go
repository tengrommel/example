package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello, world"
	//hello := s[:5]
	//world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len)
	fmt.Printf("%#v\n", []byte("hello, 世界"))
	stringChina := "中国"
	for _, item := range stringChina {
		fmt.Printf("%T\n", item)
	}
}
