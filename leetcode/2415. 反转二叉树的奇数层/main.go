package reverse_odd_levels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root} // create queue and push root
	isOdd := 0
	for len(q) > 0 {
		if isOdd != 0 {
			reverse(q)
		}
		// get the next level
		tmp := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			if node.Left!= nil { // complete binary tree (完全二叉树)
				tmp = append(tmp, node.Left, node.Right)
			}
		}
		q = tmp
		isOdd ^= 1
	}
	return root
}

func reverse(q []*TreeNode) {
	n := len(q)
	// for i, j := 0, n-1; i < j; i,j = i+1, j-1 {
	// 	q[i].Val, q[j].Val = q[j].Val, q[i].Val
	// }
	for i := 0; i < n/2; i++ {
		q[i].Val, q[n-i-1].Val = q[n-i-1].Val, q[i].Val
	}
}