package is_cousins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 使用 map 和 pair
func isCousins(root *TreeNode, x int, y int) bool {
	type pair struct {
		depth  int
		parent *TreeNode
	}
	st := make(map[int]pair)
	var dfs func(*TreeNode, *TreeNode, int) bool
	dfs = func(root, parent *TreeNode, depth int) bool {
		if root == nil {
			if p, ok := st[x]; ok {
				if q, ok := st[y]; ok {
					return p.depth == q.depth && p.parent != q.parent
				}
			}
			return false
		}

		st[root.Val] = pair{depth, parent}
		st[root.Val] = pair{depth, parent}

		if dfs(root.Left, root, depth+1) {
			return true
		}
		// if dfs(root.Right, root, depth+1) {
		// 	return true
		// }
		// return false
		return dfs(root.Right, root, depth+1)
	}
	return dfs(root, nil, 0)
}

// 使用变量记录x,y 出现的深度, 父节点 以及 是否找到
func isCousins2(root *TreeNode, x int, y int) bool {
	var dx, dy int       // depth
	var px, py *TreeNode // parent
	var fx, fy bool      // found
	var dfs func(*TreeNode, *TreeNode, int)
	dfs = func(root, parent *TreeNode, depth int) {
		if root == nil {
			return
		}
		if root.Val == x {
			dx, px, fx = depth, parent, true
		}
		if root.Val == y {
			dy, py, fy = depth, parent, true
		}
		if fx && fy {
			return
		}
		dfs(root.Left, root, depth+1)
		if fx && fy {
			return
		}
		dfs(root.Right, root, depth+1)
	}
	dfs(root, nil, 0)
	return dx == dy && px != py
}

func isCousins3(root *TreeNode, x int, y int) bool {
	var dx, dy int       // depth
	var px, py *TreeNode // parent
	var dfs func(*TreeNode, *TreeNode, int, int) (*TreeNode, int)
	dfs = func(node, parent *TreeNode, val, depth int) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		if node.Val == val {
			return parent, depth
		}
		if leftParent, depth := dfs(node.Left, node, val, depth+1); leftParent != nil {
			return leftParent, depth
		}
		// if RightParent, depth := dfs(node.Right, node, val, depth+1); RightParent != nil {
		// 	return RightParent, depth
		// }
		// return dfs(node.Right, node, val, depth+1)
		return dfs(node.Right, node, val, depth+1)
	}
	px, dx = dfs(root, nil, x, 0)
	py, dy = dfs(root, nil, y, 0)
	return dx == dy && px != py
}
