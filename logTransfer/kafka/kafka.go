package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"logTransfer/es"
)

type LogData struct {
	data string `json:"data"`
}

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
				lg := LogData{
					data: string(msg.Value),
				}
				fmt.Println("%v", string(msg.Value))
				es.SendToES(topic, lg)
			}
		}(pc)
	}
	select {}
	return nil
}
