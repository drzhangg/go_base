package main

import (
	"fmt"
	"math/rand"
	"time"
)
var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var (
	LowerCharacter = []rune("abcdefghijklmnopqrstuvwxyz")
	UpCharacter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	IntCharacter = []rune("0123456789")
)

func main() {
	rand.Seed(time.Now().UnixNano())

	result := []rune{}

	randIntCharacter :=IntCharacter[rand.Intn(len(IntCharacter))]
	randLowerCharacter :=LowerCharacter[rand.Intn(len(LowerCharacter))]
	randUpCharacter :=UpCharacter[rand.Intn(len(UpCharacter))]

	result = append(result, randIntCharacter,randLowerCharacter,randUpCharacter)

	for i := 0; i < 5; i++ {
		result = append(result,defaultLetters[rand.Intn(len(defaultLetters))])
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i],result[j] = result[j],result[i]
	})
	fmt.Println(string(result))


	// 强密码(必须包含大小写字母和数字的组合，不能使用特殊字符，长度在8-10之间)
	//match ,_ := regexp.MatchString(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,10}$`,"2VBr3WX")
	//match ,_ := regexp.MatchString(`^[a-zA-Z]\w{5,17}$`,"s2dDweqzd")
	//fmt.Println(match)

	//for i := range 5 {
	//
	//}
	//


	//rand.Shuffle()





	//b := make([]rune, 8)
	//for i := range b {
	//	b[i] = defaultLetters[rand.Intn(len(defaultLetters))]
	//}
	//
	//fmt.Println(string(b))

	//fmt.Println(RandomString(8))
}

func RandomString(n int, allowedChars ...[]rune) string {
	rand.Seed(time.Now().UnixNano())
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
