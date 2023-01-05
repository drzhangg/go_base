package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type User struct {
	Name string
	Age int
}

func main() {

	client,err := rpc.Dial("tcp","localhost:1234")
	if err != nil {
		log.Fatal("dial err:",err)
	}

	var reply User
	err = client.Call("HelloService.Hello","jerry",&reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
