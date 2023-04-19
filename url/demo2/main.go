package main

import (
	"fmt"
	"net/url"
)

func main() {
	u,err := url.Parse("abc")
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("u:",u)
}
