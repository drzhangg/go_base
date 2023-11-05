package main

import "fmt"

func main() {

	arr := []int{1,3,5,2,7,4,6}


	left,right := 0,len(arr)-1

	for left < right{
		mid := (left+right) /2

		if arr[left] >arr[mid] {

		}
			//arr[left]
	}








	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[j] > arr[i]{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}

	fmt.Println(arr)
}
