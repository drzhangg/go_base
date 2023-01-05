package main

import (
	"fmt"
)

/*
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

 */
// 旋转指针，题解 https://leetcode.cn/problems/rotate-array/solution/man-hua-san-ci-xuan-zhuan-de-fang-fa-shi-ru-he-x-2/
func rotate(nums []int, k int)  {
	k %= len(nums)   // k = k % len(nums)
	fmt.Println("k:",k)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}

func reverse(a []int) {

	for i := 0; i < len(a)/2; i++ {
		a[i],a[len(a) - 1 - i] = a[len(a) - 1 - i],a[i]
	}
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	k := 3

	for i := len(nums) -1; i >= 0 ; i-- {
		fmt.Println(nums[i])
	}

	rotate(nums,k)

	//l := len(nums)
	//
	//arr1 := nums[:l-k]
	//arr2 := nums[l-k:]
	//
	//fmt.Println(arr1)
	//fmt.Println(arr2)
}
