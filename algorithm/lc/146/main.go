package main

type LRUCache struct {
	size int
	capacity int
	cache map[int]*NodeList
	head,tail *NodeList
}

type NodeList struct {
	pre, next *NodeList
	key, val  int
}

func initNodeList(key, value int) *NodeList {
	return &NodeList{
		key: key,
		val: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache: map[int]*NodeList{},
		head:     initNodeList(0,0),
		tail:     initNodeList(0,0),
	}
	l.head.next = l.tail
	l.head.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

func main() {

}
