package main


// 两数之和
func main() {

}

/*
输入：nums = [2,7,11,15], target = 9
输出：[0,1]

输入：nums = [3,2,4], target = 6
输出：[1,2]

输入：nums = [3,3], target = 6
输出：[0,1]
 */
func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for k,v := range nums{
		if val,ok := m[target-v] ;ok{
			return []int{val,k}
		}
		m[v] = k
	}

	return nil
}
