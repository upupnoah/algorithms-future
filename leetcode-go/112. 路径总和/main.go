package _12__路径总和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自上而下
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
