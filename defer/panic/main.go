package main

import "fmt"

func defer_call() {
	defer func() {
		fmt.Println("defer: panic 之前1")
	}()

	defer func() {
		fmt.Println("defer: panic 之前2，捕获异常")

		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()

	panic("异常内容")

	defer func() {fmt.Println("defer: panic 之后, 永远执行不到") }()
}

func main() {
	defer_call()
	fmt.Println("main结束")
}
