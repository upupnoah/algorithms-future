package _603__收集树中金币

// 思路：拓扑排序
// 一个节点的度数为 1：叶子节点
// 叶子节点若是没有金币，则没有必要移动到该节点上，因此删除
// 通过先将没有金币的叶子节点 加入 队列，然后每次将队头 x 取出，删除 x 以及 让所有邻居度数-1
// 如果一个邻居的度数减少为 1 且没有金币，则继续加入队列

// 做完上面的步骤之后，现在所有的叶子节点都有金币
// 又因为：收集距离当前节点 2 以内的所有金币，因此可以移动到 叶子节点的爷爷节点（父亲的父亲）
// 操作：删除当前叶子节点，并删除新出现的叶子节点，剩余节点就是必须要访问的节点

// 题目要求返回出发点：无论从哪个点出发，每条边都必须走两次
// 答案为 剩余边数*2
// 特殊情况：最后只剩两个点，这两个点都是叶子节点，因此每个点的度数都会被删除两次，会等于-1
//
//	最后答案需要和 0 取一个最大值 max(0, ans)
func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // build graph
		deg[x]++
		deg[y]++ // 统计每个点的邻居个数 （度）
	}

	leftEdges := n - 1 // 剩余边数
	// 拓扑排序，去掉没有金币的子树
	var q []int
	for i, d := range deg {
		// 将所有没有金币的叶子节点加入队列
		if d == 1 && coins[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		leftEdges-- // 删除节点 x 到其父亲的边
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 && coins[y] == 0 {
				q = append(q, y)
			}
		}
	}

	// 再次拓扑排序
	for i, d := range deg {
		// 因为只是删了 deg，并没有真正删除这个节点
		if d == 1 && coins[i] > 0 { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来
			q = append(q, i)
		}
	}
	leftEdges -= len(q)
	for _, x := range q {
		for _, y := range g[x] {
			deg[y]--
			if deg[y] == 1 { // y现在是叶子
				leftEdges-- // 删除 y （到其父节点的边）
			}
		}
	}
	return max(leftEdges*2, 0)
}
