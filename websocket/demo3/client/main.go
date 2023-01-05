package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	dialer := websocket.Dialer{}

	connect, _ ,err := dialer.Dial("ws://127.0.0.1:8787/",nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer connect.Close()

	for {
		//从 websocket 中读取数据
		//messageType 消息类型，websocket 标准
		//messageData 消息数据
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage://文本数据
			fmt.Println("client receive:",string(messageData))
		case websocket.BinaryMessage://二进制数据
			fmt.Println(messageData)
		case websocket.CloseMessage://关闭
		case websocket.PingMessage://Ping
		case websocket.PongMessage://Pong
		default:

		}
	}
}
