package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：递归递归，先递后归
// 在向下递归的同时，额外维护一个参数 mx 表示从根节点到当前节点之前，路径上的最大值
// 边界：如果当前节点为空，返回 0
// 子问题：
//    递归左子树，获取左子树的好节点个数
//    递归右子树，获取右子树的好节点个数
// 如果当前节点是好节点 （mx <= root.Val), 返回 left + right + 1, 否则返回 left + right

func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, mx int) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left, max(mx, root.Val))
		right := dfs(root.Right, max(mx, root.Val))
		if root.Val >= mx {
			return left + right + 1
		}
		return left + right
	}
	return dfs(root, root.Val)
}

// 2023.8.26 自己再次实现
func goodNodes2(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, mx int) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left, max(mx, root.Val))
		right := dfs(root.Right, max(mx, root.Val))
		if mx <= root.Val {
			return left + right + 1
		}
		return left + right
	}
	return dfs(root, root.Val)
}
