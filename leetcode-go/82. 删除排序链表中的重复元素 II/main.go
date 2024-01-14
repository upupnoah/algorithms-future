package delete_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		if head.Val != head.Next.Val {
			pre = head
		} else {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			pre.Next = head.Next
		}
		head = head.Next
	}
	return dummy.Next
}
