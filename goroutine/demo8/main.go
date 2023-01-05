package main

import (
	"fmt"
	"sync"
	"time"
)

func job(i int) int {
	time.Sleep(time.Millisecond * 500)
	return i
}

func main() {
	num := 5
	start := time.Now()

	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(param int) {
			defer wg.Done()
			ch <- job(param)
		}(i)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()


	for c := range ch{
		fmt.Println("收到结果：",c)
	}


	end := time.Since(start)
	fmt.Println("耗时：",end.String())
}
