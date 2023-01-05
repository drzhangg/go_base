package main

import "fmt"

type Cmd func([]int) chan int

type PipeCmd func(chan int) chan int

func Evens(list []int) chan int {

	ret := make(chan int)
	go func() {
		defer close(ret)
		for _, v := range list {
			if v%2 == 0 {
				ret <- v
			}
		}
	}()

	return ret
}

func M10(in chan int) chan int {
	r := make(chan int)

	go func() {
		defer close(r)
		for ch := range in {
			r <- ch * 10
		}
	}()

	return r
}

func Pipe(args []int, c1 Cmd, c2 PipeCmd) chan int {
	ret := c1(args)
	return c2(ret)
}

func main() {
	nums := []int{2, 3, 6, 12, 22, 5, 9, 4, 64}

	ret := Pipe(nums, Evens, M10)

	for r := range ret {
		fmt.Print(" ",r)
	}

}
