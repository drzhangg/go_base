package main

import (
	"fmt"
)

func main() {

	list := []string{"123","321","abc","sdada","qweqeqw"}

	for i := 0; i < 100000; i++ {
		list = append(list, fmt.Sprintf("hello %v",i))
	}

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}


}
