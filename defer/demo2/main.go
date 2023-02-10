package main

import "fmt"

func defer1() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func defer2() (i int) {
	defer func() {
		i++
	}()
	return i
}

func main() {
	fmt.Println(defer1())
	fmt.Println(defer2())
}
