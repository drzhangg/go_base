package main

import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Name string
}

func main() {
	chanShop := make(chan Product,5)
	chanCount := make(chan int,5)
	go producer(chanShop)
	go consumer(chanShop,chanCount)

	for i := 0; i < 10; i++ {
		 <- chanCount
	}
	fmt.Println("main over")
}

// 生产者写入数据
func producer(chanShop chan<- Product) {
	for  {
		product := Product{Name: "产品"+strconv.Itoa(time.Now().Second())}
		chanShop <- product
		fmt.Println("生产者输送了",product)
		time.Sleep(time.Second)
	}
}

// 消费者消费数据，并记录消费
func consumer(chanShop <-chan Product,chanCount chan <- int) {
	for{
		product := <- chanShop
		fmt.Println("消费者消费了",product)
		chanCount <- 1
	}
}
