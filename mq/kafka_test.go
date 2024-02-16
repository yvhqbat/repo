package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"testing"
)

const (
	TopicTest = "test"
)

var kafkaAddrs = []string{"127.0.0.1:9092"}

func producer() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = TopicTest
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer(kafkaAddrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return err
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}

func consumer() error {
	consumer, err := sarama.NewConsumer(kafkaAddrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(TopicTest) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(TopicTest, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	return nil
}

func TestKafka(t *testing.T) {
	err := producer()
	if err != nil {
		return
	}

	err = consumer()
	if err != nil {
		return
	}
}
