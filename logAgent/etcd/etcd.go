package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

// 需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志要发往Kafka中的Topic
}

// 初始化ETCD的方法
func Init(addr string, timeOut time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeOut,
	})
	// watch操作
	// watch用来获取未来更改的通知
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// 从etcd根据key获取配置项
func GetConf(key string) (logEntries []*LogEntry, err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return nil, err
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntries)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed, err: %v\n", err)
			return
		}
	}
	return
}
