package main

import "fmt"

type a interface {
}

type b struct {
	name string
}

func (b *b)test() {

	b.name = "Daas"
	fmt.Println(b.name)
}

func Mul() {
	var bb  b
	bb.name = "test"

	bb.test()
	fmt.Println(bb.name)
}

func main() {
	Mul()
}
