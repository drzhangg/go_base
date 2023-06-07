package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	errMsg := `<div class=error><img src='/static/7abb227e/images/none.gif' height=16 width=1>A job already exists with the name ‘demo1’</div>`

	// 定义正则表达式模式
	pattern := `>[^<]+<`

	// 编译正则表达式
	reg := regexp.MustCompile(pattern)

	// 提取错误消息
	match := reg.FindString(errMsg)
	if match != "" {
		errorMessage := match[1:len(match)-1]
		fmt.Println(errorMessage)
	} else {
		fmt.Println("Failed to extract error message")
	}

	s2 := strings.Join([]string{"a","b"},"/job/")
	fmt.Println("/job/"+s2)
}
