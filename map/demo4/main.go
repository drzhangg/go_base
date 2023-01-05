package main

import "fmt"

type test struct {
	label map[string]string
	name string
}

func main() {

	t := test{
		label: nil,
		name:  "jerry",
	}

	if t.label == nil{
		t.label = make(map[string]string)
	}
	l := newT(&t)
	fmt.Println(l)
}

func newT(t *test) map[string]string {
	l := t.label
	l["address"] = t.name


	return l
}
