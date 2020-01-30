package main

import (
	"awesomeProject/logAgent/conf"
	"awesomeProject/logAgent/etcd"
	"awesomeProject/logAgent/kafka"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

//func run() {
//	// 读日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			kafka.SendToKafka(cfg.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//	// 发送到kafka
//}

func main() {
	err := ini.MapTo(cfg, "logAgent/conf/config.ini")
	if err != nil {
		fmt.Printf("load ini fail")
		return
	}
	// 1、初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2、打开日志文件准备收集日志
	//err = taillog.Init(cfg.FileName)
	//if err != nil {
	//	fmt.Printf("Init tailing failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("init tail log success")
	//run()
	// 2、初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	fmt.Println("init etcd success.")
	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf("teng")
	// 2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现热启动）
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value: %v\n", index, value)
	}
}
