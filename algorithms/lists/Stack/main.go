package main

import (
	"awesomeProject/algorithms/lists/Stack/StackArray"
	"fmt"
)

func FAB(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}

func main() {
	myStack := StackArray.NewStack()
	myStack.Push(7)
	last := 0
	for !myStack.IsEmpty() {
		data := myStack.Pop() // 取出数据
		if data == 1 || data == 2 {
			last += 1
		} else {
			myStack.Push(data.(int) - 1)
			myStack.Push(data.(int) - 2)
		}
	}
	fmt.Println(last)
}
