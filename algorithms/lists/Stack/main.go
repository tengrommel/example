package main

import (
	"awesomeProject/algorithms/lists/Stack/StackArray"
	"fmt"
)

func main() {
	myStack := StackArray.NewStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	stack := StackArray.NewStack()
	stack.Push(5)
	last := 0
	for !stack.IsEmpty() {
		data := stack.Pop()
		if data == nil {
			break
		}
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			stack.Push(data.(int) - 1)
		}
	}
	fmt.Println(last)
}
