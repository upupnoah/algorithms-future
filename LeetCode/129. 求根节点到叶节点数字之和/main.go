package _29__求根节点到叶节点数字之和

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 写法 1
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, sum := 0, 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			fmt.Println(sum)
			ans += sum
		}
		dfs(root.Left)
		dfs(root.Right)
		sum = (sum - root.Val) / 10
	}
	dfs(root)
	return ans
}

// 写法 2
func sumNumbers2(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += sum
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
	dfs(root, 0)
	return ans
}
