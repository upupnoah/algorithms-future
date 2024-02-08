package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x { // 答案在左子树
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > x && q.Val > x { // 答案在右子树
		return lowestCommonAncestor(root.Right, p, q)
	}
	// p < x && q > x 或者 p > x && q < x, 那么他们分别在 root 两侧, 因此当前节点 root 就是 lca
	return root
}
