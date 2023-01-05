package main

import (
	"fmt"
	"strings"
)

func main() {
	// u001b[1;32mhellou001b[m  u001b[1;34mlogsu001b[m
	a := "u001b[1;32mhellou001b[m  u001b[1;34mlogsu001b[m"

	b := strings.ReplaceAll(a,"u001b","%c")
	
	ss := []interface{}{}
	for i := 0; i < strings.Count(a,"u001b"); i++ {
		ss = append(ss, 0x1B)
	}
	fmt.Printf(b,ss...)

}
