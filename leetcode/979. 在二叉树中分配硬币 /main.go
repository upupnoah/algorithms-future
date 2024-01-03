package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：不要关注硬币的流向，而是从每条边的贡献考虑
// 对于子树 a 来说，如果子树硬币总数 x > 节点数 y，那么需要移动 x-y 个硬币
// 子树 a 的根节点到他的父亲节点的那条边上的贡献就是 x-y
// 反之，如果子树硬币总数 x < 节点数 y，那么需要从这条边上移动 y-x 个硬币过来
// 因此：不管什么情况，子树 a 的根节点到他父亲节点的那条边的贡献为 |x-y|

func distributeCoins(root *TreeNode) int {
	return dfs(root)[2]
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0, 0}
	}
	left, right := dfs(root.Left), dfs(root.Right)
	// 统计当前子树的总的 节点数和金币数
	x, y := left[0]+right[0]+1, left[1]+right[1]+root.Val
	return []int{x, y, abs(x-y) + left[2] + right[2]}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
