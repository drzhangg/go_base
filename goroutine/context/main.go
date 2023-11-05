package main

import (
	"context"
	"fmt"
	"time"
)

func cancelByContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Done()是监听cancel()、超时操作
			fmt.Println("cancel goroutine by context!")
			return
		default:
			fmt.Println("Im alive")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel1 := context.WithCancel(context.Background())
	go cancelByContext(ctx)
	time.Sleep(10 * time.Second)
	cancel1()
	time.Sleep(5 * time.Second)
}
