package main

import "fmt"

type Iterator struct {
	data []int
	index int // 索引
}

func NewIterator(data []int) *Iterator {
	return &Iterator{data: data}
}

func (i *Iterator) HasNext() bool {
	if i.data == nil || len(i.data) == 0 {
		return false
	}
	return i.index < len(i.data)
}

func (i *Iterator) Next() int {
	defer func() {
		i.index++
	}()
	return i.data[i.index]
}

func main() {
	arr := []int{1,2,3,4,5}
	iter := NewIterator(arr)

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
