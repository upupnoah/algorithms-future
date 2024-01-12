package _37__打家劫舍_III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：考虑u，a, b
// 选 u -> a与 b 都不能选，因此是 u.val + aNotRob + bNotRob
// 不选 u -> a 与 b 都有两种选择（选/不选）， max(aRob，aNotRob） + max(bRob, bNotRob)
// 答案为 root 选/不选 的最大值
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	// 返回两个状态，当前这一层的子树 选/不选
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0 // 如果是空节点，那么选这个节点 和 不选这个节点都是一样的(0)
		}
		lRob, lNotRob := dfs(root.Left)
		rRob, rNotRob := dfs(root.Right)
		// 选当前节点
		rob := lNotRob + rNotRob + root.Val
		notRob := max(lRob, lNotRob) + max(rRob, rNotRob)
		return rob, notRob
	}
	rob, notRob := dfs(root)
	return max(rob, notRob)
}
