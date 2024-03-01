package main

// DFS
func reachableNodes1(n int, edges [][]int, restricted []int) int {
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	st := make(map[int]bool)
	for _, v := range restricted {
		st[v] = true
	}
	// visited := make(map[int]bool) // 使用一个map来记录访问过的节点
	// 可以使用一个 visited 数组标记是否访问过, 也可以在 dfs 中多传入一个参数, 表示当前节点的 father
	var dfs func(int, int) int
	dfs = func(u int, fa int) int {
		if st[u] {
			return 0
		}
		// visited[u] = true
		res := 1 // 当前节点本身也算一个可到达的节点
		for _, out := range g[u] {
			// if !visited[out] {}
			if out != fa { // 检查是否访问过，避免重复计算
				res += dfs(out, u)
			}
		}
		return res
	}
	return dfs(0, -1)
}

// BFS
func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	st := make(map[int]bool)
	for _, v := range restricted {
		st[v] = true
	}
	visited := make(map[int]bool)
	q := make([]int, 0)
	q = append(q, 0)
	ans := 1
	for len(q) > 0 {
		var t []int
		for _, v := range q {
			visited[v] = true
			for _, node := range g[v] {
				if !visited[node] && !st[node] {
					ans++
					t = append(t, node)
				}
			}
		}
		q = t
	}
	return ans
}
