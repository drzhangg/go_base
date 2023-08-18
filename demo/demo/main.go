package main

import "fmt"

type ListNode struct {
	Val  string
	Next *ListNode
}

func (l *ListNode) Add(add *ListNode) *ListNode {
	// add header
	add.Next = l
	return add
}

func main() {
	var m map[string]struct{}

	val := m["name"]
	fmt.Println(val)
}
