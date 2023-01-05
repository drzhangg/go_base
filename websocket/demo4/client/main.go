package main

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"net/http"
)

func main() {
	//u := "http://127.0.0.1:8887/echo"
	//wsu := "ws://127.0.0.1:8887/echo"
	//ws,err := websocket.Dial(u,"",wsu)
	//if err != nil {
	//	fmt.Println("err::", err)
	//}
	//
	//defer ws.Close()
	//
	//for  {
	//	data := make([]byte, 1024*10)
	//	_, err = ws.Read(data)
	//	if err != nil {
	//		log.Println("err:", err)
	//	}
	//
	//	fmt.Println("data::", string(data))
	//}

	sockjs.NewHandler("/",sockjs.DefaultOptions, func(session sockjs.Session) {

	})

	http.ListenAndServe("",nil)


}

func test() {
	//dialer := websocket.Dialer{}
	//
	//connect, _, err := dialer.Dial("ws://127.0.0.1:8887/echo", nil)
	//if nil != err {
	//	log.Println("dial err:",err)
	//	return
	//}
	//defer connect.Close()
	//
	//for {
	//	messageType, messageData, err := connect.ReadMessage()
	//	if nil != err {
	//		log.Println(err)
	//		break
	//	}
	//	fmt.Println("type::", messageType)
	//	fmt.Println("data::", messageData)
	//}
}
