package main

import "fmt"

// 思路
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// 双指针，左指针从数组下标0开始
	left := 0

	// 右指针直接从数组的下标1开始，
	for i := 1; i < len(nums); i++ {
		// 左右指针对应的下标值进行比较，如果右指针值和左指针值不一样，那就把左指针的位置加1，并且把右指针对应的值给左指针
		if nums[left] != nums[i]{
			left++
			nums[left] = nums[i]
		}
	}
	// 最后因为返回的是长度，所以需要把左指针的位置加1返回
	left++
	return left
}

func main() {
	nums:= []int{1,2,2,3,3,3,4}
	fmt.Println(removeDuplicates(nums))
}
