package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 生产者等待leader和follower都回复Ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true // not clear

	msg := &sarama.ProducerMessage{}
	msg.Topic = "zyl"
	msg.Value = sarama.StringEncoder("zyl love zsj")

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil{
		fmt.Println(err)
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("%v offset %v", pid, offset)

}
