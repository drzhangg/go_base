package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "aaaaaa123ereraaaaaaa"
	b := "a"
	c := strings.Trim(a, b)
	fmt.Println(c)

	d := strings.Split(a, b)
	fmt.Println(d)
}
