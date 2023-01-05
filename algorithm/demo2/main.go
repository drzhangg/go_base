package main

import "math"

// 中级
// 给一个自然数数组l，再给一个自然数n，找出数组l里连续的k个数使得它们的和大于等于n，输出这个k的最小值，如果不存在输出0。
func main() {

}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	start,end := 0,0
	sum :=0

	for end < n{
		sum += nums[end]
		for sum >= target{
			ans = min(ans,end - start + 1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if ans == math.MaxInt32{
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x > y{
		return y
	}
	return x
}
