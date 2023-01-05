package main

import "fmt"

// 中级
// 实现一个函数，入参是 从1开始的整数排序数组 l, 翻转系数数组k_list 。
// K翻转的规则如下：翻转系数数组的元素 k：代表要翻转的数字个数，第一次从第1个数往后数k个进行翻转，
// 得到的结果再从第2个数往后数k个进行翻转，直到翻转到最后的k个数，表示完成了一次k的翻转。
// k_list中每一个元素都要对l进行k翻转。返回翻转之后的数组。
func main() {
	// []int{1,3,4,5,7,9,10}   3
	// []int{}
	//arr := []int{1,3,4,5,7,9,10}
	//k_list := []int{1,3,4,5,7,9,10}

	//na := reverse(arr)
	//fmt.Println("new:",na)

	arr := []int{1,2,3,4,5}
	k_list := []int{1,2,3,4,5}

	fmt.Println(fx(arr,k_list))


}

/*
def fx2(l, k_list, x):
    n = len(l)
    for k in k_list:
        for i in range(0, n - k + 1):
            l = l[:i] + list(reversed(l[i: i + k])) + l[i + k:]
    return l[x - 1]
 */


/*
[]{1,2,3,4,5}   k=2
2,3,4,5,1
 */




func fx(l, k_list []int)[]int {
	n := len(l)
	result1 := []int{}
	for _, k := range k_list {
		if (n-k+1)%2 ==0{
			result := []int{}
			result = append(result, l[k-1:]...)
			result = append(result, l[:k-1]...)
			result1 = result
		}else {
			result := []int{}
			result = append(result, l[k-1:]...)
			na := reverse(l[:k-1])
			result = append(result, na...)
			result1 = result
		}
	}
	return result1
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
