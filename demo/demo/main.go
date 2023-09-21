package main

import (
	"fmt"
	"sync"
)

func letterPrinter(wg *sync.WaitGroup, letter string, ch chan bool, nextCh chan bool) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ch // 等待上一个goroutine完成
		fmt.Print(letter)
		nextCh <- true // 发送信号给下一个goroutine
	}
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	wg.Add(3)

	go letterPrinter(&wg, "H", ch1, ch2)
	go letterPrinter(&wg, "O", ch2, ch3)
	go letterPrinter(&wg, "H", ch3, ch1)

	// 开始第一个goroutine
	ch1 <- true

	wg.Wait()
	fmt.Println() // 打印换行符，以便下一行输出
}
