package main

import "fmt"

func main() {
	n := 115

	fmt.Println(n%10)

	n1 := n >> 2
	fmt.Println(n1)

	num := 3
	k := 3

	k %= num

	fmt.Println(k %num)
	fmt.Println(k)
}

