package main

import "fmt"

func main() {
	//nums := []int{1,2,3,4,5,6,7,7,9,10,12,12,14}
	nums := []int{1,2,3,4,5,6,7,7,9}
	fmt.Println(splitArray(nums,4))
}

func splitArray(array []int, n int) [][]int {
	length := len(array)
	quotient := length / n
	remainder := length % n
	fmt.Println(quotient)
	fmt.Println(remainder)

	result := make([][]int, n)
	currentIndex := 0

	for i := 0; i < n; i++ {
		size := quotient
		if i < remainder {
			size++
		}

		result[i] = make([]int, size)
		copy(result[i], array[currentIndex:currentIndex+size])
		currentIndex += size
	}

	return result
}
