package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

func Init(addr []string, maxSize int) (err error) {
	config := sarama.NewConfig()

	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer error", err)
		return err
	}

	logDataChan = make(chan *logData, maxSize)
	go sentToKafka()
	return nil
}

func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

func sentToKafka() error {
	for {
		select {
		case lg := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = lg.topic
			msg.Value = sarama.StringEncoder(lg.data)

			_, _, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send message to kafka error, ", err)
				return err
			}

			//fmt.Printf("pid %v offset %v\n", pid, offset)
		}

	}
}
