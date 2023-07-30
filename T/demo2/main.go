package main

import "fmt"

type ID interface {
	int | string
}

type IdData[id ID] struct {
	Id id
}

func (i IdData[id])GetId()  {
	fmt.Println(i.Id)
	fmt.Printf("%T\n",i.Id)
}

func main() {
	var id1 = IdData[int]{1}
	var id2 = IdData[string]{"123"}

	id1.GetId()
	id2.GetId()
}
