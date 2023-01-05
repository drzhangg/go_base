package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func sum(nums []int,target int) [][]int {
	result := [][]int{}
	if len(nums) == 0 {
		return nil
	}

	left,right := 0,len(nums) -1

	for left < right {
		if nums[left] + nums[right] > target{
			right--
		}else if nums[left] + nums[right] < target {
			left++
		}else {
			result = append(result, []int{nums[left],nums[right]})
			left++
			right--
		}
	}
	return result
}


func sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i+1; j < len(nums); j++ {

			if nums[j] < nums[minIdx]{
				minIdx = j
			}
		}
		nums[i],nums[minIdx] = nums[minIdx],nums[i]
	}
}

//func sort(nums []int) []int {
//
//	n := len(nums)
//	for i := 1; i < n; i++ {
//		tmp := nums[i]
//		j := i -1
//		for j >=0&& nums[j] > tmp{   //左边比右边大
//			nums[j+1] = nums[j] // 右移1位
//			j--   //扫描前一个数
//		}
//		nums[j+1] = tmp
//	}
//	return nums
//}


func main() {
	v := 123_453
	fmt.Println(v)

	arr := []int{1,3,2,6,4,5}
	sort(arr)
	fmt.Println(arr)

	result := sum(arr,8)
	fmt.Println(result)
}
