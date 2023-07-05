package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 双指针：先走 k 步，然后一起走
// a | a | a | a | a | nil
//           | a |   k=2
// 要求倒数第二个节点，可以让一个指针先走 k=2步，然后一起走
// 当先走的指针到达 nil（终点）的时候，后走的指针正好走到倒数第二个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	for k != 0 {
		p2 = p2.Next
		k--
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}