package main

import (
	"awesomeProject/logAgent/log-transfer/conf"
	"awesomeProject/logAgent/log-transfer/es"
	"awesomeProject/logAgent/log-transfer/kafka"
	"fmt"
	"gopkg.in/ini.v1"
)

// log transfer
// 将日志数据从kafka取出来发往ES
func main() {
	// 加载配置文件
	var cfg = new(conf.LogTransfer) // 返回指针
	err := ini.MapTo(cfg, "logAgent/log-transfer/conf/cfg.ini")
	// 初始化
	if err != nil {
		fmt.Printf("init config, err:%v\n", err)
		return
	}
	fmt.Printf("cfg: %v\n", cfg)
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Printf("es conected err:%v \n", err)
		return
	}

	// 1、从kafka取日志数据
	// 连接kafka创建分区的消费者
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	// 2、发往ES
	if err != nil {
		fmt.Printf("init kafka consumer failed, err: %v\n", err)
		return
	}

}
