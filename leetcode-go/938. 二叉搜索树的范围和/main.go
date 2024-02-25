package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func rangeSumBST1(root *TreeNode, low int, high int) int {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Val > high {
			return dfs(root.Left)
		}
		if root.Val < low {
			dfs(root.Right)
		}
		return root.Val + dfs(root.Left) + dfs(root.Right)
	}
	return dfs(root)
}

func rangeSumBST2(root *TreeNode, low int, high int) (ans int) {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 利用 BST 的特性, 只遍历某一边
		if root.Val > high {
			dfs(root.Left)
			return
		}
		if root.Val < low {
			dfs(root.Right)
			return
		}
		// 两边都需要遍历(上面的条件一个都没满足)
		ans += root.Val
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return
}

// BFS
func rangeSumBST3(root *TreeNode, low int, high int) (ans int) {
	var q []*TreeNode
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		if node.Val < low {
			q = append(q, node.Right)
			continue
		}
		if node.Val > high {
			q = append(q, node.Left)
			continue
		}
		ans += node.Val
		q = append(q, node.Left, node.Right)
	}
	return
}
