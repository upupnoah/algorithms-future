package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution1_me: 迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre = head.Next
	head.Next = nil
	for pre != nil {
		tmp := pre.Next
		pre.Next = head 
		head = pre
		pre = tmp
	}
	return head
}

// solution1_official: 迭代
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
