package main

import (
	"awesomeProject/logAgent/conf"
	"awesomeProject/logAgent/kafka"
	"awesomeProject/logAgent/taillog"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 读日志
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(cfg.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
	// 发送到kafka
}

func main() {
	err := ini.MapTo(cfg, "logAgent/conf/config.ini")
	if err != nil {
		fmt.Printf("load ini fail")
		return
	}
	// 1、初始化kafka连接
	// 2、打开日志文件准备收集日志
	err = kafka.Init([]string{cfg.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	err = taillog.Init(cfg.FileName)
	if err != nil {
		fmt.Printf("Init tailing failed, err:%v\n", err)
		return
	}
	fmt.Println("init tail log success")
	run()
}
