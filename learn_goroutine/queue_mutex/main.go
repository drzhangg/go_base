package main

import "sync"

type SliceQueue struct {
	data []interface{}
	sync.Mutex
}

func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{data: make([]interface{},0,n)}
}

// Enqueue 把值放在队尾
func (s *SliceQueue) Enqueue(v interface{})  {
	s.Lock()
	s.data = append(s.data, v)
	s.Unlock()
}

// Dequeue 移去队头并返回(出队)
func (s *SliceQueue) Dequeue()  interface{}{
	s.Lock()
	if len(s.data) == 0 {
		s.Unlock()
		return nil
	}
	v := s.data[0]
	s.data = s.data[1:]
	s.Unlock()
	return v
}

func main() {

}
