package main

import (
	"awesomeProject/logAgent/kafka"
	"awesomeProject/logAgent/taillog"
	"fmt"
	"time"
)

func run() {
	// 读日志
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
	// 发送到kafka
}

func main() {
	// 1、初始化kafka连接
	// 2、打开日志文件准备收集日志
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init Kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init tailing failed, err:%v\n", err)
		return
	}
	fmt.Println("init tail log success")
	run()
}
