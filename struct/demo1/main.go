package main

import (
	"fmt"
)

type EventType string

const (
	Added EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted EventType = "DELETED"
)

type Event struct {
	Type EventType
}

func main() {
	e := Event{}

	//reflect.Map


	fmt.Println(e)
}
