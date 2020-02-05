package main

import (
	"awesomeProject/algorithms/lists/CircleQueue"
	"fmt"
)

func main() {
	var myQueue CircleQueue.CircleQueue
	CircleQueue.InitQueue(&myQueue)
	CircleQueue.EnQueue(&myQueue, 1)
	CircleQueue.EnQueue(&myQueue, 2)
	CircleQueue.EnQueue(&myQueue, 3)
	CircleQueue.EnQueue(&myQueue, 4)
	CircleQueue.EnQueue(&myQueue, 5)
	fmt.Println(CircleQueue.DeQueue(&myQueue))
	fmt.Println(CircleQueue.DeQueue(&myQueue))
	fmt.Println(CircleQueue.DeQueue(&myQueue))
	fmt.Println(CircleQueue.DeQueue(&myQueue))
	fmt.Println(CircleQueue.DeQueue(&myQueue))
}
