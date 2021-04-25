package etcd

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

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
