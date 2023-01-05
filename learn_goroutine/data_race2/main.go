package main

import (
	"fmt"
	"sync"
)

func main() {
	var(
		counter = 0
		wg sync.WaitGroup
	)

	wg.Add(10)

	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()

			for i := 0; i < 100000; i++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}