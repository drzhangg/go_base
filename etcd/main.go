package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

//var (
//	nodeId = flag.Int("id",0,"node ID")
//	addr = flag.String("addr","http://150.158.87.137:2379","etcd address")
//
//)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://150.158.87.137:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	ctx := context.Background()

	getResp,err := cli.Get(ctx,"name")
	if err != nil {
		fmt.Printf("get to etcd failed, err:%v\n", err)
		return
	}

	for _, v := range getResp.Kvs{
		fmt.Println(v.String())
	}

	// put
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//_, err = cli.Put(ctx, "name", "dsb")
	//cancel()
	//if err != nil {
	//	fmt.Printf("put to etcd failed, err:%v\n", err)
	//	return
	//}


	//var (
	//	config.conf  clientv3.Config
	//	err     error
	//	client  *clientv3.Client
	//)
	//
	//ctx := context.Background()
	//
	////配置
	//config.conf = clientv3.Config{
	//	Endpoints:   []string{"150.158.87.137:2379"},
	//	DialTimeout: time.Second * 5,
	//}
	//
	////连接 创建一个客户端
	//if client, err = clientv3.New(config.conf); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////用于读写etcd的键值对
	////kv = clientv3.NewKV(client)
	//
	//_,err = client.Put(ctx,"name","jerry")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//getResp, err = kv.Get(ctx, "/cron/jobs/job1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}
