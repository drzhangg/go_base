package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 先定义一个虚拟节点
	dump := &ListNode{}
	dump.Next = head

	// 定义快慢指针，快指针先走n+1步，然后快慢指针一起走
	slow,fast := dump,dump

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}


	// 直到快指针走到等于null，删除慢指针的next节点
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	delNode := slow.Next
	slow.Next = slow.Next.Next
	delNode.Next = nil

	return dump.Next
}

func main() {

}
