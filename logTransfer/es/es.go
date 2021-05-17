package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func Init(address string) (err error) {
	client, err = elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		fmt.Println("es init error, ", err)
		return err
	}
	fmt.Println("connect to es successfully")
	return nil
}

func SendToES(index string, data interface{}) (err error) {
	put1, err := client.Index().Index(index).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		fmt.Println("es put msg error ", err)
		return err
	}

	fmt.Printf("%s has been sent to es %s\n", put1.Index, put1.Id)
	return nil
}
