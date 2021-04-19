package main

import (
	"fmt"
	"github.com/doudou215/LogCollection/logAgent/kafka"
	"github.com/doudou215/LogCollection/logAgent/tailLog"
	"gopkg.in/ini.v1"
	"logAgent/conf"
	"time"
)

var cfg = new(conf.AppConf)

func run() {
	for {
		select {
		case line := <-tailLog.GetTailChan():
			kafka.SentToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//cfg, err := ini.Load("./conf/config.ini")
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("ini load error ", err)
		return
	}

	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("kafka init error ", err)
		return
	}

	err = tailLog.Init(cfg.TailLogConf.Filename)
	if err != nil {
		fmt.Println("tail init error ", err)
		return
	}

	run()
}
