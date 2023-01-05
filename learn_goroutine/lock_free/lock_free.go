package lock_free

import (
	"sync/atomic"
	"unsafe"
)

type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{
		head: n,
		tail: n,
	}
}

// 将unsafe.Pointer原子加载转换成node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装CAS，避免直接将*node转换成unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v}

	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) {
			if next == nil {
				if cas(&tail.next, next, n) {
					cas(&q.tail, tail, n)
					return
				}
			} else {
				cas(&q.tail, tail, next)
			}
		}
	}
}

func (q *LKQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) {
			// head还是那个head
			if head == tail { // head和tail一样
				if next == nil { // 说明是空队列
					return nil
				} // 只是尾指针还没有调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else { // 读取出队的数据
				v := next.value // 既然要出队了，头指针移动到下一个
				if cas(&q.head, head, next) {
					return v // Dequeue is done.
				}
			}
		}
	}
}
