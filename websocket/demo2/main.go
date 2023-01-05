package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {

	wsurl := "ws:"

	connect,_,err :=websocket.DefaultDialer.Dial(wsurl,nil)
	if err != nil {
		log.Println("err:", err)
		return
	}
	defer connect.Close()

	for {
		_, messageData, err := connect.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}


		//fmt.Println("type::", messageType)
		fmt.Println("time:",time.Now(),"data::", string(messageData))
	}
}
