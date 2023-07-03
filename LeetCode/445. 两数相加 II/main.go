package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 存入栈中，然后用 t 模拟进位
// 头插法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := []int{}, []int{}
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var head *ListNode
	n, m := len(s1), len(s2)
	t := 0
	for n > 0 || m > 0 || t != 0 {
		if n > 0 {
			t += s1[n-1]
			n--
		}
		if m > 0 {
			t += s2[m-1]
			m--
		}
		// 头插法
		node := &ListNode{Val: t % 10}
		node.Next = head
		head = node
		t /= 10
	}
	return head
}