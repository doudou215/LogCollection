package main

import (
	"fmt"
	"github.com/doudou215/LogCollection/logAgent/etcd"
	"github.com/doudou215/LogCollection/logAgent/kafka"
	"github.com/doudou215/LogCollection/logAgent/tailLog"
	"gopkg.in/ini.v1"
	"logAgent/conf"
	"time"
)

var cfg = new(conf.AppConf)

/*
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
*/

func main() {
	//cfg, err := ini.Load("./conf/config.ini")
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("ini load error ", err)
		return
	}

	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.MaxChanSize)
	if err != nil {
		fmt.Println("kafka init error ", err)
		return
	}

	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)

	//fmt.Println("key ", cfg.EtcdConf.Key)
	logEntries, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		return
	}

	for _, ev := range logEntries {
		fmt.Printf("key: %v value：%v\n", ev.Path, ev.Topic)
	}

	tailLog.Init(logEntries)
	newConfChan := tailLog.GetNewConfChan()
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	select {}
	/*
		err = tailLog.Init(cfg.TailLogConf.Filename)
		if err != nil {
			fmt.Println("tail init error ", err)
			return
		}
	*/
	// run()
}
