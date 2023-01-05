package main

import (
	"fmt"
)


type A int

func main() {
	m := make(map[A]string)
	a := A(1)
	b := A(2)
	c := A(3)
	m[a] = "a"
	m[b] = "b"
	m[c] = "c"
	fmt.Println(m)
}
