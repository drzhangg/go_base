package main

import (
	"fmt"
	"go_base/interface/demo2"
)

func main() {
	c := test()
	apps := c.AppsV1()
	fmt.Println(apps)

	core := c.CoreV1()
	fmt.Println(core)
}

func test() demo2.Interface {
	client := demo2.NewClient()
	return client
}
