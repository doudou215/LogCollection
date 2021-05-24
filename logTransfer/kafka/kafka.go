package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"logTransfer/es"
)

func Init(address, topic string) (err error) {
	consumer, err := sarama.NewConsumer([]string{address}, nil)
	if err != nil {
		fmt.Println("kafka init error, ", err)
		return err
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Partitions error ", err)
		return err
	}
	fmt.Printf("partitionList %v address %s topic %s \n", partitionList, address, topic)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer pc.AsyncClose()
		go func(partitionConsumer sarama.PartitionConsumer) {
			fmt.Println("waiting for message")
			for msg := range pc.Messages() {
				lg := es.LogData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				fmt.Printf("%s\n", string(msg.Value))
				es.SendToESchan(&lg)
			}
		}(pc)
	}
	select {}
	return nil
}
