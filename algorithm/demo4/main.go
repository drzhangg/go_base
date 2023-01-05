package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5,6}
	k_list := []int{1, 2, 3, 4, 5,6}

	k := 2

	fmt.Println(fx(arr,k_list, k))
}

func fx1(arr []int, k int) []int {


	start, end := 0, k
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

func fx2(arr, k_list []int) []int {
	n := len(arr)

	if k_list[len(k_list)-1] > n {
		return nil
	}


	for i := 0; i < len(k_list); i++ {
		if (n-k_list[i]+1)%2 ==0 {

		}
	}

	return arr
}


func fx(arr, k_list []int, k int) []int {
	if k_list[len(k_list)-1] > len(arr) {
		return nil
	}

	for i := 0; i < len(k_list); i++ {
		start, end := 0, k_list[i]
		for end <= len(arr)  {

			r2 := []int{}
			r2 = append(r2, arr[:start]...)
			r2 = append(r2, reverse(arr[start:end])...)
			r2 = append(r2, arr[end:]...)
			arr = r2
			start++
			end++
		}
	}
	return arr
}

func reverse(arr []int) []int {
	arrlen := len(arr)
	temp := 0

	for i := 0; i < arrlen/2; i++ {
		temp = arr[arrlen-1-i]
		arr[arrlen-1-i] = arr[i]
		arr[i] = temp
	}
	return arr
}
