package main

func search(nums []int, target int) int {
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] == target {
	//		return i
	//	}
	//}

	start,end := 0,len(nums)-1
	for start <= end {
		mid := start + (end- start) /2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			end = mid - 1
		}else {
			start = mid + 1
		}
	}

	return -1
}

func main() {

}
