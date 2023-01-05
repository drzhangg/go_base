package main

import (
	"fmt"
	"sync"
)

func main() {

	list := []string{"ads","qweqw","weq"}

	ch := make(chan struct{},10)

	wg := sync.WaitGroup{}

	for _,v := range list{
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			fmt.Println(v)

			ch <- struct{}{}
		}(v)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for{
		select {
		case _,ok := <-ch:
			if !ok {
				return
			}
		}
	}
}
