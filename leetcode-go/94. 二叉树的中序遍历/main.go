package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var dfs func(*TreeNode)
	var ans []int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)

	}
	dfs(root)
	return ans
}