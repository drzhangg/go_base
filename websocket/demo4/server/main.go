package main

import (
	"fmt"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"net/http"
	"time"
)

func main() {

	http.Handle("/",sockjs.NewHandler("/",sockjs.DefaultOptions, func(session sockjs.Session) {
		fmt.Println("进入handler")
		for {
			err := session.Send("send from server by sockjs")
			if err !=nil{
				fmt.Println("send err:",err)
				break
			}
			time.Sleep(time.Second)
		}
	}))
	//handler := sockjs.NewHandler("/", sockjs.DefaultOptions, echoHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}


