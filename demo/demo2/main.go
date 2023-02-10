package main

import "fmt"

type Config struct {
	
}

func (c *Config) NewConfig(name, ip string, Role []string) *Config{
	return nil
}

func main() {

	numbers := []int{1,3,2,6,4,5}
	//fmt.Println(sort(nums))

	//numbers := []int{1, 2, 3, 4, 5}
	target := 8
	result := findPairsWithSum(numbers, target)
	fmt.Println(result)
}

func allSum() [][]int {
	return nil
}

func findPairsWithSum(numbers []int, target int) [][]int {
	result := [][]int{}
	// 创建哈希表，用于存储第一个整数
	m := make(map[int]bool)
	for i := 0; i < len(numbers); i++ {
		complement := target - numbers[i]
		// 判断是否存在第二个整数，使得它们的和为指定值
		if m[complement] {
			result = append(result, []int{complement, numbers[i]})
		}
		m[numbers[i]] = true
	}
	return result
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
