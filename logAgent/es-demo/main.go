package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (s *Student) run() {
	fmt.Printf("%s在跑...", s.Name)
}

func (s *Student) wang() {
	fmt.Printf("%s汪汪汪地叫...", s.Name)
}

func main() {

	luminghui := Student{
		Name:    "卢明辉",
		Age:     34,
		Married: false,
	}
	luminghui.run()
	luminghui.wang()

	// 1、初始化连接，得到一个client
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := Student{Name: "rion", Age: 44, Married: false}
	// 链式操作
	put1, err := client.Index().
		Index("student").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
