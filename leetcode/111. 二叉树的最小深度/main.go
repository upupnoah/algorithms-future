package _11__二叉树的最小深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：递归
// 考虑 f(u) 与 f(a) f(b)的关系
// 1. u是叶子，return 1
// 2. u 不是叶子
//   2.1 a b 都不空  return min(f(a), f(b))+1
//   2.2 a 不空 return f(a)+1
//   2.3 b 不空 return f(b)+1

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左边空返回 f(b)
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	// 右边空，返回 f(a)
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	// 左右都不空,返回 min(f(a), f(b))
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 思路 2：bfs
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q []*TreeNode
	q = append(q, root)
	ans := 1
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			t := q[0]
			q = q[1:]
			if t.Right == nil && t.Left == nil {
				return ans
			}
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		ans++
	}
	return ans
}
