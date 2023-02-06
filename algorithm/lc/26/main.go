package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

/*
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]

题解：
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solutions/34033/shuang-zhi-zhen-shan-chu-zhong-fu-xiang-dai-you-hu/
*/
func removeDuplicates(nums []int) int {

	// 双指针
	q, p := 0, 1
	for p < len(nums) {
		if nums[q] != nums[p] {
			nums[q+1] = nums[p]
			q++
		}
		p++
	}
	return q + 1

}
