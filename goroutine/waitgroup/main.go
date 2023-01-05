package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	s1 := []string{"1","2","3","4","5","6"}

	var s2 []string
	for i := 0; i < 6; i++ {
		s2 = append(s2, strconv.Itoa(i))
	}

	wg := sync.WaitGroup{}

	for _,v := range s1{

		fmt.Println("v:",v)

		for _,v1 := range s2{
			wg.Add(1)

			go func(v1 string) {
				defer wg.Done()

				fmt.Println("v1:",v1)
			}(v1)
		}
	}

	wg.Wait()

}
