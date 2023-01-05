package main

import "fmt"


func main() {
	//fmt.Println("main:", test())

	arr := []int{0,1,2,3}
	m := make(map[int]*int)

	for k,v := range arr{
		m[k] = &v
	}

	for _,v := range m{
		fmt.Println(*v)
	}
}

func test() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

