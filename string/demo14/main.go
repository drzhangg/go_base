package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""

	arr := strings.Split(s,",")
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(len(arr))

	arr1 := []string{"",""}
	fmt.Println(arr1)
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
}
