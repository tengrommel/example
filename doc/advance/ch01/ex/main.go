package main

import "fmt"

func main() {
	array := [...]int{
		1, 2: 3,
	}
	fmt.Println(len(array))
	for index := range array {
		fmt.Println(array[index])
	}
	b := &array
	for index, item := range b {
		fmt.Println(index, item)
	}
	var times [12][2]int
	var number int
	for range times {
		number++
	}
	fmt.Println(number)

	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<-c2
}
