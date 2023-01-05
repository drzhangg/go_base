package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4 // 要使用的goroutine数量
	taskLoad = 10 //要处理的工作的数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建一个有缓冲的通道来管理工作
	tasks := make(chan string,numberGoroutines)

	wg.Add(numberGoroutines)

	for i := 1; i <= numberGoroutines; i++ {
		go  worker(tasks,i)
	}

	for i := 0; i <= taskLoad; i++ {
		tasks <- fmt.Sprintf("Task : %d",i)
	}

	// 当所有工作都处理完时关闭通道,以便所有goroutine退出
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for{
		task,ok := <- tasks
		if !ok{
			//意味着通道已经关闭
			fmt.Printf("Worker: %d : Shutting Down\n",worker)
			return
		}

		// 显示开始工作
		fmt.Printf("Worker: %d : Started %s\n",worker,task)

		// 随机等一段时间来模拟工作
		sleep := rand.Intn(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n",worker,task)
	}
}
