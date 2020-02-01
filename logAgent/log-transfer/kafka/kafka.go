package kafka

import (
	"awesomeProject/logAgent/log-transfer/es"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

// LogData ...
type LogData struct {
	data string `json:"data"`
}

// 初始化kafka连接的一个client 准备好kafka消费者
func Init(addrs []string, topic string) error {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err: %v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err: %v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d Offset: %d Key: %v Value: %v",
					msg.Partition, msg.Offset, msg.Key, msg.Key)
				var ld = new(LogData)
				err := json.Unmarshal(ld, msg.Value)
				if err != nil {
					fmt.Printf("unmarshal failed. err: %v\n", err)
					return
				}
				es.SendToES(topic)
			}
		}(pc)
	}
	return nil
}
