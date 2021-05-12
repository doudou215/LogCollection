package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logTransfer/conf"
	"logTransfer/es"
	"logTransfer/kafka"
)

func main() {
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/config.ini")
	if err != nil {
		panic(err)
	}

	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Println("es init error")
		return
	}

	err = kafka.Init(cfg.KafkaCfg.Address, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Println("kafka init error")
		return
	}

	select {}
}
