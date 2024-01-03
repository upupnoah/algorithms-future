package _17__填充每个节点的下一个右侧节点指针_II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// BFS
// 时间复杂度：O(n)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		t := q // 当前这一层的所有节点
		q = nil
		for i, node := range t {
			if i < len(t)-1 {
				node.Next = t[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		t[len(t)-1].Next = nil
	}
	return root
}
