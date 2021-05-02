package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

var cli *clientv3.Client

func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Println("etcd initialize error ", err)
	}
	fmt.Println("connected to etcd successfully")
	return err
}

// 传进来的参数是从ini中读取到的存放在etcd中的（key, value)的key
// 已知key的情况下从etcd中读取value，这些value是通过jason的形式存放的
// 所以要先经过反序列化才能用
func GetConf(key string) (logEntries []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ret, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get configuration error ", err)
		return nil, err
	}
	for _, ev := range ret.Kvs {
		// 反序列化
		//fmt.Println(ev.Value)
		err = json.Unmarshal(ev.Value, &logEntries)
		if err != nil {
			fmt.Println("json decodes fail ", err)
			break
		}
	}
	return logEntries, err
}
