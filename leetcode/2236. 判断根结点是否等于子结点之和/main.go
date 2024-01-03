package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Right.Val+root.Left.Val == root.Val
}

// 拓展到 n 个子节点怎么做？ 递归
func checkTreeN(root *TreeNode) bool {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return root.Val + dfs(root.Left) + dfs(root.Right)
	}
	return dfs(root.Left)+dfs(root.Right) == root.Val
}
