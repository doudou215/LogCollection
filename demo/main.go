package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("client initialize error ", err)
		return
	}

	fmt.Println("connected to etcd successfully")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"D:/apache/kafka_2.13-2.7.0/logs/server.log","topic":"server_log"}]`
	_, err = cli.Put(ctx, "log2topic", value)
	cancel()
	if err != nil {
		fmt.Println("etcd put error ", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	rets, err := cli.Get(ctx, "log2topic")
	cancel()
	if err != nil {
		fmt.Println("etcd get error", err)
		return
	}

	for _, ev := range rets.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	/*
		ch := cli.Watch(context.Background(), "zyl")
		for ret := range ch {
			for _, ev := range ret.Events {
				fmt.Printf("type %s, key %s, value %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	*/
}
