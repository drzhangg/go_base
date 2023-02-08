package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

//func keys[K comparable, V any](m map[K]V) []K {
//
//}

type All interface {
	string | int | map[string]string | interface{} | chan int
}

type Slice1[any int | float64 | string] []any

type Map1[KEY int | string, VALUE string | float64] map[KEY]VALUE

type MapTest[key int | string, value string | interface{} | user | int] map[key]value

type user struct {
	Name string
	Age  int
}

type Struct1[any string | int] struct {
	Title   string
	Content any
}

type struct2[any string|int, any2 Slice1[any]] struct {
	Name any
	Title any2
}

func Sum[T int | float64](a, b T) T {
	return a + b
}

type testStruct[any1 string | int, any2 map[any1]any3,any3 struct{}|interface{}] struct {
	Name    string
	Content any1
	Job     any2
}

func max[T constraints.Ordered] (a,b T) T {
	if a > b {
		return a
	}
	return b
}

func match[t comparable](a, b t) bool {
	return a==b
}


func main() {


	//fmt.Println(match(1,2))
	//fmt.Println(match(1.45,1.45))
	//fmt.Println(match("string","string"))
	//fmt.Println(match(true,true))

	var age int = 1
	var sex int = 1
	p1 := &age
	p2 := &sex
	fmt.Println(match(p1, p2))




	fmt.Println(max(1,2))
	fmt.Println(max(1.233,1.22))
	fmt.Println(max("hello", "small"))

	s2 := struct2[string,Slice1[string]]{
		Name: "zhang",
		Title: []string{
			"123","456","heheheheh",
		},
	}
	fmt.Println(s2)

	tt := testStruct[string,map[string]struct{},struct{}]{
		Name: "jerry",
		Job: map[string]struct{}{"123": {}},
		Content: "hello",
	}
	fmt.Println(tt)

	var stuct1 Struct1[string]
	stuct1.Title = "this is title"
	stuct1.Content = "this is content"

	struct2 := Struct1[int]{
		Title:   "hehhehehe",
		Content: 2023,
	}
	fmt.Println(stuct1.Content, struct2.Content)

	m1 := Map1[string, string]{
		"1": "zhang",
		"2": "jerry",
	}
	fmt.Println(m1["2"])

	m2 := MapTest[int, user]{
		1: user{},
		2: user{
			Age:  21,
			Name: "jerry",
		},
	}
	fmt.Println(m2[2].Name)

	fmt.Println(Sum[int](1, 2))
	fmt.Println(Sum[float64](1.23, 2.54))

	var MySlice1 Slice1[int] = []int{1, 2, 3}

	fmt.Println(MySlice1)

	mys2 := Slice1[string]{"hello", "world"}
	fmt.Println(mys2)
}
