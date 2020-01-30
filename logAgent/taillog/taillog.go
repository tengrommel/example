package taillog

import (
	"awesomeProject/logAgent/kafka"
	"fmt"
	"github.com/hpcloud/tail"
)

// TailTask: 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tail *TailTask) {
	tail = &TailTask{
		path:  path,
		topic: topic,
	}
	tail.init() // 根据路径打开对应的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}
	go t.run() // 直接去收集采集日志
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			// 3.2 发往Kafka 将同步的调用变成异步
			kafka.SendToKafka(t.topic, line.Text)
		}
	}
}
