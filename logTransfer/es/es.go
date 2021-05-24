package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type LogData struct {
	Data  string `json:"data"`
	Topic string `json:"topic"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

func Init(address string) (err error) {
	client, err = elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		fmt.Println("es init error, ", err)
		return err
	}
	fmt.Println("connect to es successfully")
	ch = make(chan *LogData, 10000)
	go SendToES()
	return nil
}

func SendToESchan(msg *LogData) {
	ch <- msg
}

func SendToES() (err error) {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Println("es put msg error ", err)
				return err
			}
			fmt.Printf("%s has been sent to es %s\n", put1.Index, put1.Id)
		default:
			time.Sleep(time.Second)
		}
	}

	return nil
}
