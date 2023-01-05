package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func job() error{
	ctx,_ := context.WithTimeout(context.Background(),time.Second)

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Millisecond * 500)
		done <- struct{}{}
	}()

	select {

	case <-done:
		return nil

	case <-ctx.Done():
		return fmt.Errorf("超时了")
	}
}

func main() {
	//fmt.Println(job())

	for i := 0; i < 20; i++ {
		go func() {
			job()
		}()
	}

	for{
		time.Sleep(time.Second * 2)
		fmt.Println(runtime.NumGoroutine())
	}
}
