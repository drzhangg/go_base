package main

func twoSum(numbers []int, target int) []int {
	//m := make(map[int]struct{})
	l,r := 0,len(numbers) -1

	for numbers[l] + numbers[r] != target {
		if numbers[l] + numbers[r] > target{
			r--
		}else {
			l++
		}
	}
	return []int{l+1,r+1}
}

func main() {

}
