package main

import (
	"fmt"
	"sync"
)

func main() {
	list := []string{"123","321","abc","sdada","qweqeqw"}

	wg := sync.WaitGroup{}

	ch := make(chan struct{},4)
	defer close(ch)

	for _,v := range list{
		ch <- struct{}{}
		wg.Add(1)

		go func(v string) {
			defer wg.Done()

			fmt.Println("v:",v)
			<-ch
		}(v)
	}
	wg.Wait()
}
