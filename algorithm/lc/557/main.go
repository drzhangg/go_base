package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	arrs := strings.Split(s," ")

	var result []string
	for _,v := range arrs{

		result = append(result, fmt.Sprintf("%v",reverse([]byte(v))))
	}

	return strings.Join(result," ")
}

func reverse(s []byte) string {
	for l,r := 0,len(s) -1;l<r;l++{
		s[l],s[r] = s[r],s[l]
		r--
	}
	return string(s)
}

func main() {
	//s := "God Ding"
	s := "Let's take LeetCode contest"
	r := reverseWords(s)
	fmt.Println(r)
}
