package main

import (
	"awesomeProject/algorithms/lists/Queue"
	"fmt"
)

func main() {
	myQueue := Queue.NewQueue()
	myQueue.EnQueue(1)
	myQueue.EnQueue(2)
	myQueue.EnQueue(3)
	myQueue.EnQueue(4)
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.DeQueue())
}
