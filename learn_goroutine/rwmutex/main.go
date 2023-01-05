package main

import (
	"sync"
	"time"
)

func main() {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			for{
				counter.Count() // 计数器读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for{ // 一个writer
		counter.Incr()  // 计数器写操作
		time.Sleep(time.Second)
	}
}

type Counter struct {
	mu sync.RWMutex
	count int64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
