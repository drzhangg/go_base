package main

func findRepeatNumber1(nums []int) int {
	result := 0
	m := make(map[int]struct{})
	for k,v := range nums{

		_,ok := m[v]
		if ok{
			result = nums[k]
		}
		m[v] = struct{}{}
	}

	return result
}

func findRepeatNumber(nums []int) int {
	l,r :=0, nums[len(nums) - 1]

	m := make(map[int]int)
	result := 0
	for l <= r {
		_,ok := m[r]
		if ok {
			result = nums[r]
		}

		_,ok = m[l]
		if ok {
			result = nums[l]
		}

		m[nums[l]] = nums[l]
		m[nums[r]] = nums[r]
		l++
		r--
	}
	return result
}

func main() {

}
