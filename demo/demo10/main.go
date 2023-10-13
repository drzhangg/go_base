package main

import "fmt"

type MyStruct struct {
	Val int
}

func (d *MyStruct) Modify(val int) {
	d.Val = val
}

func (d MyStruct) Modify1(val int) {
	d.Val = val
}

func main() {

	d := MyStruct{Val: 2}

	d.Modify(4)
	fmt.Println(d)

	d.Modify1(3)
	fmt.Println(d)

}
