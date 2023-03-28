package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3
	rate := time.Tick(time.Second / 3)
	bucket := make(chan int, n)
	go func() {
		for {
			bucket <- 1
		}
	}()
	for i := 0; i < 10; i++ {
		<-bucket
		<-rate

		go printt(i)
		//go func(req int) {
		//	fmt.Printf("Processing request %d\n", req)
		//}(i)
	}
}

func printt(req int) {
	fmt.Printf("Processing request %d\n", req)
}
