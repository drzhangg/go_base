package main

import "fmt"

func main() {
	numbers := []int{1,3,2,6,4,5}
	fmt.Println(findSlice(numbers,6))
}

func findSlice(arr []int, target int) [][]int {
	result := [][]int{}

	m := make(map[int]bool) // 默认bool是false

	for i := 0; i < len(arr); i++ {
		other := target - arr[i]

		if m[other]{
			result = append(result, []int{other,arr[i]})
		}
		m[arr[i]] = true
	}
	return result
}
