package main

import (
	"context"
	"log"
	"time"
)

func g(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("g取消了")
		return
	}
}

func main() {
	ctx,cancel := context.WithCancel(context.Background())

	go g(ctx)

	count := 1
	for{
		if count >=3{
			cancel()
		}
		count++
		time.Sleep(time.Second)
	}
}
