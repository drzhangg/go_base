package main

func moveZeroes(nums []int)  {

	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i],nums[i+1:]...)
			count++
			i--
		}
	}

	for i := 0; i < count; i++ {
		nums = append(nums, 0)
	}
}

func main() {

}
