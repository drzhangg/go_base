package main

import "fmt"

func partition(list []int, low, high int) int {

	pivot := list[high]

	for low < high{

		for low < high && pivot >= list[low]{
			low++
		}

		list[high] = list[low]

		for low < high &&pivot <= list[high]{
			high--
		}

		list[low] = list[high]
	}

	list[high]=pivot
	return high
}

func QuickSort(list []int,low, high int)  {
	if high > low{
		pivot := partition(list,low,high)

		QuickSort(list,low,pivot-1)

		QuickSort(list,pivot+1,high)
	}
}

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
