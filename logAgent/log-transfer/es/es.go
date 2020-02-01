package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

var (
	client *elastic.Client
)

// 初始化es，准备接收kafka发来的数据
// Init ...
func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success")
	return nil
}

// 发送数据到ES
func SendToES(indexStr, typeStr string, data interface{}) (err error) {
	put1, err := client.Index().Index(indexStr).BodyJson(data).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
