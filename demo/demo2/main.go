package main

import "fmt"

func main() {
	nums := []int{1,3,2,6,4,5,3}
	fmt.Println(sort(nums))
}

func allSum() [][]int {
	return nil
}

// 排序
func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {

			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}
