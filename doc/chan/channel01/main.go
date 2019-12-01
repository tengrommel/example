package main

import "fmt"

func main() {
	// 定义一个ch1变量
	// 是一个channel类型
	// 这个channel内部传递的数据是int类型
	var ch1 chan int
	var ch2 chan string
	// channel 是引用类型
	fmt.Println("ch1:", ch1)
	fmt.Println("ch2:", ch2)
	// make函数初始化(相当于分配内存)： slice map channel
	ch3 := make(chan int)
	// 通道的操作：发送、接收、关闭
	// 发送和接收 <-
	ch3 <- 10 // 把10发送到ch3中

}
