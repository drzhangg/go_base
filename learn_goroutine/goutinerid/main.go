package main

import "fmt"

func main() {
	d := []int{1,2,3,5,2,6,6,8}
	r := uniq(d)
	fmt.Println(r)
}

func uniq(data []int) []int {
	r := []int{}
	m := make(map[int]struct{})

	for _, v := range data{
		m[v] = struct{}{}
	}

	for k,_ := range m{
		r = append(r, k)
	}



	return r
}
