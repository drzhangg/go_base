package main

import "fmt"

func main() {
	arr := []int{1,2,3}
	arr2 := arr[:]

	copy(arr2,arr)

	arr[1] = 0
	fmt.Println(arr)
	fmt.Println(arr2)
}
