package main

import "fmt"

func main() {
  nums:= []int{3,2,4}
  target := 6

  fmt.Println(twoSum(nums,target))
}


/*
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
 */
func twoSum(nums []int, target int) []int {

	//m := make(map[int]struct{})

	//for i:=0;i < len(nums);i++{
	//	for j := i+1; j < len(nums); j++ {
	//		if nums[i] + nums[j] == target{
	//			return []int{i,j}
	//		}
	//	}
	//}


	 m := make(map[int]int)

	 for k,v := range nums{

		 if _,ok := m[target - v];ok{
			 return []int{m[target-v],k}
		 }
		 m[v] = k

	 }
	return []int{}
}