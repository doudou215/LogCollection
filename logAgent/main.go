package main

import (
	"github.com/doudou215/LogCollection/logAgent/kafka"
	"github.com/doudou215/LogCollection/logAgent/tailLog"
	"time"
)

func run() {
	for {
		select {
		case line := <- tailLog.GetTailChan():
			kafka.SentToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		print(err)
		return
	}

	err = tailLog.Init("./myLog")
	if err != nil {
		print(err)
		return
	}
	
	go run()
}
