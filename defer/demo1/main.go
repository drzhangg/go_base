package main

import "fmt"

func main() {

	var whatever [3]struct{}

	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
