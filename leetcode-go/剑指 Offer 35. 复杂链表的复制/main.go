package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 方法 1：回溯+哈希表
var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
    if node == nil {
        return nil
    }
    if n, has := cachedNode[node]; has {
        return n
    }
    newNode := &Node{Val: node.Val}
    cachedNode[node] = newNode
    newNode.Next = deepCopy(node.Next)
    newNode.Random = deepCopy(node.Random)
    return newNode
}

func copyRandomList(head *Node) *Node {
    cachedNode = map[*Node]*Node{}
    return deepCopy(head)
}

// 方法 2：迭代 + 节点拆分
func copyRandomList2(head *Node) *Node {
    if head == nil {
        return nil
    }
    for node := head; node != nil; node = node.Next.Next {
        node.Next = &Node{Val: node.Val, Next: node.Next}
    }
    for node := head; node != nil; node = node.Next.Next {
        if node.Random != nil {
            node.Next.Random = node.Random.Next
        }
    }
    headNew := head.Next
    for node := head; node != nil; node = node.Next {
        nodeNew := node.Next
        node.Next = node.Next.Next
        if nodeNew.Next != nil {
            nodeNew.Next = nodeNew.Next.Next
        }
    }
    return headNew
}