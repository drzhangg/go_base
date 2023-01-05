package main

import "fmt"

func main() {
	a:="A"
	b:="A"
	c:= &a==&b

	fmt.Println(c)
}
