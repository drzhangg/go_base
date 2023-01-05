package main

import (
	"fmt"
	"strings"
)

var (
	romanMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"a": 4,
		"b": 9,
		"c": 40,
		"d": 90,
		"e": 400,
		"f": 900,
	}
)

func romanToInt(s string) int {
	s = strings.Replace(s, "IV", "a", -1)
	s = strings.Replace(s, "IX", "b", -1)
	s = strings.Replace(s, "XL", "c", -1)
	s = strings.Replace(s, "XC", "d", -1)
	s = strings.Replace(s, "CD", "e", -1)
	s = strings.Replace(s, "CM", "f", -1)
	var result int
	for i := 0; i < len(s); i++ {
		v := fmt.Sprintf("%c", s[i])
		val, ok := romanMap[v]
		if ok {
			result += val
		}
	}
	return result
}

func main() {
	s := "MCMXCIV"
	r := romanToInt(s)
	fmt.Println(r)
}
