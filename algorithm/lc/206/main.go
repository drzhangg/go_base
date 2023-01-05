package main

// 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return n

}

func main() {


}
