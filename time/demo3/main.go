package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设这是你从 Kubernetes 事件中获取的时间
	eventTime := time.Now().Add(-90 * time.Minute) // 90 minutes ago

	// 计算时间差
	duration := time.Since(eventTime)

	// 格式化时间差
	//hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// 输出
	fmt.Printf("LAST SEEN: %dm%ds ago\n", minutes, seconds)
}
