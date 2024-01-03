package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建一个虚拟头结点，最后的答案是 dummy.Next
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	t := 0
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: t % 10}
		cur = cur.Next
		t /= 10
	}
	return dummy.Next
}
