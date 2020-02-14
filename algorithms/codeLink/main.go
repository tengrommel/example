package main

import "fmt"

func main() {
	myQueue := NewLinkQueue()
	for i := 0; i < 100; i++ {
		myQueue.Enqueue(i)
	}
	for data := myQueue.Dequeue(); data != nil; data = myQueue.Dequeue() {
		fmt.Println(data)
	}
}
