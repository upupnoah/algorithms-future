package insert_greatest_common_divisors

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil {
		slow.Next = &ListNode{
			Val:  gcd(slow.Val, fast.Val),
			Next: fast,
		}
		slow, fast = fast, fast.Next
	}
	return head
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
