package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2024-03-21T08:16:28Z"

	// 解析时间字符串
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Println("解析时间失败:", err)
		return
	}

	// 将时间转换为其他时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("加载时区失败:", err)
		return
	}

	localTime := t.In(loc)

	fmt.Println("UTC 时间:", t)
	fmt.Println("上海时间:", localTime)
}
