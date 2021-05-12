package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		panic(err)
	}

	plist, err := consumer.Partitions("test_log")
	if err != nil {
		panic(err)
	}
	fmt.Println("partition list ", plist)
	for partition := range plist {
		pc, err := consumer.ConsumePartition("test_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer pc.AsyncClose()
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition %d key %v value %v\n", msg.Partition, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	select {}
}
