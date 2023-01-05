package main

func main() {
	//sort.Sort()


}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k,v := range nums{

		if val,ok := m[target-v];ok {
			return []int{val,k}
		}

		m[k] = v

	}
	return nil
}
