package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:               "150.158.87.137:6379",
		Password:           "",
		DB:                 0,
	})

	logrus.Infof("")

	ctx := context.TODO()

	cmd := rdb.Get(ctx,"name")
	fmt.Println(cmd.String())

	//for i := 0; i < 1000; i++ {
	//
	//	rdb.PFAdd("codehole",fmt.Sprintf("user%s",i))
	//
	//	intCmd := rdb.PFCount("codehole")
	//	fmt.Println(intCmd.String())
	//	//if intCmd.String() != i+1 {
	//	//
	//	//}
	//}


	//inserted,err :=rdb.Do(ctx,"BF.ADD","bf_key","item0").Bool()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//if inserted{
	//	fmt.Println("item0 was inserted")
	//}else {
	//	fmt.Println("item0 already exists")
	//}

	for i := 0; i < 100000; i++ {
		rdb.Do(ctx,"bf.add","codehole",fmt.Sprintf("user%d",i))
	}



}
