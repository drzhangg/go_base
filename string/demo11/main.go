package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 20

	fmt.Println(skip(n))
}

func skip(n int) int {
	if n <4 {
		 return 0
	}

	count := 0
	for i := 0; i < n; i++ {
		sn := strconv.Itoa(i)

		if strings.Contains(sn,"4"){
			count++
		}
	}
	return count

}
