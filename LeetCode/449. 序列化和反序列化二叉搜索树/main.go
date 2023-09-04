package _49__序列化和反序列化二叉搜索树

import (
	"strconv"
	"strings"
)

// 二叉搜索树
// 思路：
// 1. 可以直接忽略 BST 这个条件，使用 BFS 或者直接当满二叉树来序列化和反序列化
// 2. 前序遍历 + BST 特性（左边的<=根，右边的严格>根），或者 左边的<根，右边的>=根（本题没有重复的元素）
// 	2.1 前序遍历进行序列化：首先对于某个子树而言，其必然是在这个序列化的 s 中连续存储（即能够使用 s[l,r]表示）
//      同时首位元素必然是该子树的头结点
//  2.2 反序列化：将 s 根据分隔符 “ ” 进行分割，假设分割后数组 ss 长度为 n，那么 ss[0,n−1] 代表完整的子树
// 		我们可以利用「二叉树」特性递归构建，设计递归函数 dfs2(l, r int, ss []string) TreeNode
//		其含义为利用 ss[l,r] 连续段构造二叉树，并返回头结点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var str []string
	var dfs1 func(*TreeNode)
	dfs1 = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 前序遍历：根左右
		str = append(str, strconv.Itoa(root.Val))
		dfs1(root.Left)
		dfs1(root.Right)
	}
	dfs1(root)
	return strings.Join(str, " ")
}

// Deserializes your encoded data to tree.
func (Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ss := strings.Split(data, " ")
	var dfs2 func(int, int, []string) *TreeNode
	dfs2 = func(l, r int, ss []string) *TreeNode {
		if l > r {
			return nil
		}
		j := l + 1
		t, _ := strconv.Atoi(ss[l])
		ans := &TreeNode{Val: t}
		for j <= r {
			v, _ := strconv.Atoi(ss[j])
			if v > t {
				break
			} else {
				j++
			}
		}
		ans.Left = dfs2(l+1, j-1, ss)
		ans.Right = dfs2(j, r, ss)
		return ans
	}
	return dfs2(0, len(ss)-1, ss)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
