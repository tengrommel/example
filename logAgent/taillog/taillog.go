package taillog

import (
	"awesomeProject/logAgent/kafka"
	"context"
	"fmt"
	"github.com/hpcloud/tail"
)

// TailTask: 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能够实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tail *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tail = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
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
	// 当goroutine执行的函数退出的时候，goroutine就结束了
	go t.run() // 直接去收集采集日志
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task: %v  结束了...", t.path+t.topic)
			return
		case line := <-t.instance.Lines:
			// 3.2 发往Kafka 将同步的调用变成异步
			//kafka.SendToKafka(t.topic, line.Text)
			// 先把日志数据发到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取日志数据发到kafka
		}
	}
}
