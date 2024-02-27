package main

import "fmt"

func main() {

	m1 := make(map[string][]string)
	m1["one"] = []string{}
	m1["two"] = []string{}
	m1["three"] = []string{}

	a1 := "one"
	//a2 := "two"
	//a3 := "three"

	//for{
	m1[a1] = append(m1[a1], "1")
	m1[a1] = append(m1[a1], "2")
	m1[a1] = append(m1[a1], "3")
	//}

	fmt.Println(m1)

}
