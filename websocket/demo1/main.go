package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func main() {
	u := "http://127.0.0.1:8989/ws/v1/container/log/"
	wsu := "ws://127.0.0.1:8989/ws/v1/container/log/689/4l31pfeq/websocket"


	ws, err := websocket.Dial(wsu, "", u)
	if err != nil {
		fmt.Println("err::", err)
	}

	defer ws.Close()

	for  {
		data := make([]byte, 1024*10)
		_, err = ws.Read(data)
		if err != nil {
			log.Println("err:", err)
		}

		fmt.Println("data::", string(data))
	}

}
