package main

import "fmt"

func findTarget(nums []int,target int) [][]int {
	result := [][]int{}
	m := make(map[int]bool)

	for _,val := range nums{
		flag := target - val
		if m[flag]{
			result = append(result, []int{val,flag})
		}
		m[val] = true
	}
	return result
}

func main() {
	nums := []int{1,2,3,5,6,8,4}
	fmt.Println(findTarget(nums,9))
}
