package main

import (
	"fmt"
	"reflect"
)

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

type Int interface {
	~int | int8 | int32
}

type Float interface {
	float32 | float64
}

type SliceElement interface {
	Int | Float
}

type Slice1[T SliceElement] []T

func main() {
	var s1 Slice1[int]

	type MyInt int

	var s2 Slice1[MyInt]

	fmt.Println(s1,s2)

	ms := MySlice[int]{
		1,2,3,4,
	}

	var m1 MySlice[float32] = []float32{1.2,2.3,4.1}

	fmt.Println(m1.Sum())

	result := ms.Sum()
	fmt.Println(result)
	fmt.Println(reflect.ValueOf(result).Kind())


	r := Add(3,5)
	fmt.Println(r)
	fmt.Println(reflect.ValueOf(r).Kind())
}
