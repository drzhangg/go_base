package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}


	// 创建客户端连接
	client,err:= elasticsearch.NewClient(cfg)
	//client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
		return
	}

	res,err := client.Ping(client.Ping.WithContext(context.TODO()))
	if err != nil {
		fmt.Printf("elasticsearch.Ping failed, err:%v\n", err)
		return
	}
	defer res.Body.Close()

	if res.IsError(){
		fmt.Println("Error:", res.StatusCode)
		return
	}
	fmt.Println("Elasticsearch is up and running!")


	// get document

	client.Get("","")


	//getReq := esapi.GetRequest{
	//	Index:          "",
	//	DocumentID:     "",
	//}

	//getRes,err := getReq.Do(context.TODO(),client)


	//client.Search(client.Search.WithIndex(""),client.Search.WithBody(""))



}
