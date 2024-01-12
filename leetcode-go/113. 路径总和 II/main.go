package _13__路径总和_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		sum -= root.Val
		if root.Left == nil && root.Right == nil && sum == 0 {
			// ❌，切片是引用类型，后续修改会更改 ans 中的 path
			//ans = append(ans, path)
			// ✅
			ans = append(ans, append([]int{}, path...))
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return ans
}
