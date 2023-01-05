package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "max+without%28endpoint%29+%28sum+without%28"
	//r,_ := url.QueryUnescape(str)
	//fmt.Println(r)

	s := "test-dev"
	arr := strings.Split(s,"-")
	if len(arr) >=2{

		ns := arr[1:]
		fmt.Println(ns)

		na := strings.Join(ns,"-")
		fmt.Println("na:",na)
	}
}
