package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	config := flag.String("config", "", "巡检配置文件")

	file, err := os.Open(*config)
	//cd, err = os.ReadFile(*config)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件出错:", err)
		return
	}

	// 处理读取的数据，这里简单地直接打印出来
	fmt.Println("content:", string(content))

}
