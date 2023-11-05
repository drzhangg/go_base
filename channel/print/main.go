package main

import (
	"fmt"
	"sync"
)

func printLetters(wg *sync.WaitGroup, letterCh chan bool, numberCh chan bool) {
	defer wg.Done()

	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {
		<-letterCh // 阻塞等待数字协程发送信号
		fmt.Print(letter)
		numberCh <- true // 发送信号给数字协程
	}

	<-letterCh // 阻塞等待数字协程发送最后一个信号
}

func printNumbers(wg *sync.WaitGroup, letterCh chan bool, numberCh chan bool) {
	defer wg.Done()

	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		<-numberCh // 阻塞等待字母协程发送信号
		fmt.Print(number)
		letterCh <- true // 发送信号给字母协程
	}

	<-numberCh // 阻塞等待字母协程发送最后一个信号
}

func main() {
	wg := &sync.WaitGroup{}
	letterCh := make(chan bool)
	numberCh := make(chan bool)


	wg.Add(1)
	go printLetters(wg, letterCh, numberCh)

	go printNumbers(wg, letterCh, numberCh)

	numberCh <- true // 启动字母协程

	wg.Wait()
	fmt.Println()



}
