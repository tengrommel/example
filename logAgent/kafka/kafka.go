package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

var client sarama.SyncProducer
var loadDataChan chan *logData

// 给外部暴露的一个函数
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	loadDataChan <- msg
}

func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return err
	}
	loadDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道中取出
	go sendToKafka()
	return
}

// 真正往Kafka发送日志的函数
func sendToKafka() {
	for {
		select {
		case msgStruct := <-loadDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = msgStruct.topic
			msg.Value = sarama.StringEncoder(msgStruct.data)
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
