package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：快慢指针 + 反转链表 + 交错合并
func reorderList(head *ListNode) {
	mid := middleNode(head)
	head2 := reverseList(mid)

	// 交错合并
	for head2.Next != nil {
		nxt := head.Next
		head.Next = head2
		nxt2 := head2.Next
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
