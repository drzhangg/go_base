package main

import "fmt"

func main() {

	if !build(){
		fmt.Println(123)
		return
	}
	fmt.Println(456)
}

func build() bool{
	a := false

	return !a && false
}
