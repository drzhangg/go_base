package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func test() {
	log.Println(" ===> loop begin")

	for i := 0; i < 1000; i++ {
		log.Println(genSomeBytes())
	}

	log.Println(" ===> loop end")
}

func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for i := 0; i < 20000; i++ {
		buff.Write([]byte{'0'+ byte(rand.Intn(10))})
	}

	return &buff
}

func main() {

	go func() {
		for{
			test()
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe(":10000",nil)
}
