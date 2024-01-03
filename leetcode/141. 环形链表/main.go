package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
