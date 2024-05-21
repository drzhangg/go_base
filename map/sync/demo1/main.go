package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Every(time.Minute*1), 1)
	fmt.Println(limiter.Allow())
	fmt.Println(limiter)
	// 创建一个新的 sync.Map
	var m sync.Map

	// 将键值对存储到 sync.Map 中
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// 从 sync.Map 中获取值
	val1, ok1 := m.Load("key1")
	if ok1 {
		fmt.Println("Value for key1:", val1)
	}

	// 使用 Range 遍历 sync.Map 中的所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true // 返回 true 继续迭代，返回 false 停止迭代
	})



	// 删除键值对
	m.Delete("key1")

	// 使用 LoadOrStore 方法尝试获取值，如果键不存在，则存储指定的值
	val2, loaded := m.LoadOrStore("key2", "new value")
	if loaded {
		fmt.Println("Value for key2:", val2, "(already existed)")
	} else {
		fmt.Println("New value stored for key2:", val2)
	}

	// 检查 sync.Map 中是否存在某个键
	_, ok2 := m.Load("key1")
	if !ok2 {
		fmt.Println("key1 does not exist in the map")
	}
}
