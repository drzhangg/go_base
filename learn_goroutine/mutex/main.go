package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	counter int64
}

func (c *Counter) Incr()  {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}
func (c *Counter) Count() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {
	var count Counter

	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()

			for i := 0; i < 100000; i++ {
				count.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.Count())
}