package main

import "fmt"

/**
每次扩充就会重新申请一块区域
*/
func main() {
	a := make([]int64, 0, 32)
	fmt.Println(cap(a), len(a))
	for i := 0; i < 30; i++ {
		a = append(a, 1)
		fmt.Println(cap(a), len(a))
	}
}
