package _123__最深叶节点的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：自下而上，递归做
// 考虑 u 和 a， b 的关系，返回 lca 和 depth
// 对于当前子问题：
//  1. a.depth = b.depth, return u, depth+1
//  2. a.depth > b.depth, return a.lca, depth+1
//  3. a.depth < b.depth, return b.lca, depth+1
//
// 答案为 dfs(root)
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (*TreeNode, int)
	dfs = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}
		left, leftD := dfs(root.Left)
		right, rightD := dfs(root.Right)
		if leftD == rightD {
			return root, leftD + 1
		}
		if leftD > rightD {
			return left, leftD + 1
		}
		return right, rightD + 1
	}
	ans, _ := dfs(root)
	return ans
}
