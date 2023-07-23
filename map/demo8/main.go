package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string)

	// 填充map
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		m[key] = value
	}

	// 并发读取数据
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := m[key]
			fmt.Printf("Read: %s -> %s\n", key, value)
		}(i)
	}

	wg.Wait()
}
