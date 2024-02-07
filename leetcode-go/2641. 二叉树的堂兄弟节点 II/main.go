package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		t := q
		q = nil
		// 计算下一层的节点值之和
		nextLevelSum := 0
		for _, node := range t {
			if node.Left != nil {
				q = append(q, node.Left)
				nextLevelSum += node.Left.Val
			}
			if node.Right != nil {
				q = append(q, node.Right)
				nextLevelSum += node.Right.Val
			}
		}

		// 再次遍历
		for _, node := range t {
			val := 0
			if node.Left != nil {
				val += node.Left.Val
			}
			if node.Right != nil {
				val += node.Right.Val
			}
			if node.Left != nil {
				node.Left.Val = nextLevelSum - val
			}
			if node.Right != nil {
				node.Right.Val = nextLevelSum - val
			}
		}
	}
	return root
}
