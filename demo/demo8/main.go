package main

import "fmt"

type I interface {
	Do()
}

type A struct {
	Success bool
}

func (a *A) Do() {
	fmt.Println(a.Success)
}

func NewA() I {
	return &A{}
}

func NewI() (i I, err error) {
	i = NewA()
	if obj, ok := i.(*A); ok {
		obj.Success = true
	}
	return
}

func main() {
	i, err := NewI()
	if err != nil {
		panic(err)
	}
	i.Do()
}
