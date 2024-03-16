package main

// 拓扑排序
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)
	degree := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
		degree[x]++
		degree[y]++
	}
	var leaves []int
	for i, d := range degree {
		// 叶子节点
		if d == 1 {
			leaves = append(leaves, i)
		}
	}
	// 一层一层剥离, 最后 n <= 2, 剩下的就是答案
	for n > 2 {
		n -= len(leaves)
		var newLeaves []int
		for _, leaf := range leaves {
			for _, neighbor := range adj[leaf] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}
		leaves = newLeaves
	}
	return leaves
}

// BFS
func findMinHeightTrees3(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}
	parents := make([]int, n)
	bfs := func(start int) (x int) {
		q := []int{start}
		vis := make([]bool, n)
		vis[start] = true
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, neighbor := range adj[x] {
				if !vis[neighbor] {
					parents[neighbor] = x
					vis[neighbor] = true
					q = append(q, neighbor)
				}
			}
		}
		return
	}
	x := bfs(0)
	y := bfs(x)
	var path []int
	parents[x] = -1
	for y != -1 { // 从 y 开始回溯到 x, 直到 x (怎么表示到了 x? parent[x] = -1)
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}

// DFS
func findMinHeightTrees4(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}
	parents := make([]int, n)
	maxDepth, node := 0, -1
	var dfs func(int, int, int)
	dfs = func(x, pa, depth int) {
		if depth > maxDepth {
			maxDepth, node = depth, x
		}
		parents[x] = pa
		for _, y := range adj[x] {
			if y != pa {
				dfs(y, x, depth+1)
			}
		}
	}
	dfs(0, -1, 1)
	maxDepth = 0
	dfs(node, -1, 1)
	path := []int{}
	for node != -1 {
		path = append(path, node)
		node = parents[node]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}
