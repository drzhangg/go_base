package main

import (
	"sort"
)

/*
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]


输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

 */
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}

func main() {
	nums := []int{-4,-1,0,3,10}

	sortedSquares(nums)
}
