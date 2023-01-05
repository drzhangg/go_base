package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	upgrader := websocket.Upgrader{}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		connect, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println("err:", err)
			return
		}
		defer connect.Close()

		go tickWriter(connect)

		for{
			messageType,messageData,err := connect.ReadMessage()
			if err !=nil{
				log.Println(err)
				break
			}

			switch messageType {
			case websocket.TextMessage://文本数据
				fmt.Println(string(messageData))
			case websocket.BinaryMessage://二进制数据
				fmt.Println(messageData)
			case websocket.CloseMessage://关闭
			case websocket.PingMessage://Ping
			case websocket.PongMessage://Pong
			default:

			}
		}

	})

	err := http.ListenAndServe(":8787", nil)
	if nil != err {
		log.Println(err)
		return
	}
}

func tickWriter(connect *websocket.Conn) {
	for {
		err := connect.WriteMessage(websocket.TextMessage, []byte("from sever to client"))
		if err != nil {
			log.Println(err)
			break
		}
		time.Sleep(time.Second)
	}
}
