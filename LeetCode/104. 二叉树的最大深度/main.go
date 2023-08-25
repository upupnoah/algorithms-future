package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 做法 1，将返回值作为答案
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}

// 做法 2，维护一个全局变量
// Leetcode 全局变量的坑，提交的时候会出错
// 共享的一个全局变量，因此推荐下面这种写法
func maxDepth2(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		ans = max(ans, depth)
		dfs(root.Left, depth)
		dfs(root.Right, depth)
	}
	dfs(root, 0)
	return ans
}
