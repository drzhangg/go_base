package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	k := 2

	fmt.Println(fx(arr,k))
}

func fx(arr []int, k int) []int {
	start,end := 0,k
	for {
		r2 := []int{}


		r2 = append(r2, arr[:start]...)
		r2 = append(r2, reverse(arr[start:end])...)
		r2 = append(r2, arr[end:]...)
		arr = r2
		start++
		end++

		if end > len(arr) {
			break
		}
	}
	return arr
}



func reverse(arr []int) []int {
	arrlen := len(arr)
	temp := 0

	for i :=0 ;i <arrlen /2;i++ {
		temp = arr[arrlen-1-i]
		arr[arrlen-1-i] = arr[i]
		arr[i] = temp
	}
	return arr
}

