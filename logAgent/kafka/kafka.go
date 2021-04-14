package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(addr []string) (err error) {
	config := sarama.NewConfig()

	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil{
		fmt.Println("producer error", err)
		return err
	}
	return nil
}

func SentToKafka(topic, context string) error {
	msg := &sarama.ProducerMessage{}
//	上面的msg为啥不能是普通变量, 小傻瓜，因为client.SendMessage需要指针类型
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(context)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message to kafka error, err")
		return err
	}

	fmt.Printf("%v read msg %v\n", pid, offset)
	return nil
}
