package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"marry"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("err ", err)
		return
	}

	fmt.Println("connect to es successfully")
	p1 := Person{Name: "zyl", Age: 18, Married: true}

	// 典型的链式调用，所以在Index返回的是client的引用
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("index %s type: %s\n", put1.Index, put1.Type)
}
