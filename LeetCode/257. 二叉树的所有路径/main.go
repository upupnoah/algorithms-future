package _57__二叉树的所有路径

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自上而下，先处理再递归
func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var dfs func(*TreeNode, string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ans = append(ans, path)
		}
		path += "->"
		dfs(root.Left, path)
		dfs(root.Right, path)
	}
	dfs(root, "")
	return ans
}

// dfs不需要第二个参数的写法
func binaryTreePaths2(root *TreeNode) []string {
	var ans []string
	var path []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			line := strconv.Itoa(path[0])
			for i := 1; i < len(path); i++ {
				line += "->" + strconv.Itoa(path[i])
			}
			ans = append(ans, line)
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return ans
}
