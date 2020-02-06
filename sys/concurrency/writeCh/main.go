package main

import (
	"fmt"
	"time"
)

func writeToChan(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

// The keyword is used for declaring that the c function parameter will
// be a channel, and it should be followed by the type of channel(int)

func main() {
	c := make(chan int)
	go writeToChan(c, 10)
	time.Sleep(time.Second)
}
