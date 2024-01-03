package main

func maximumScoreAfterOperations(edges [][]int, values []int) (ans int64) {
	n := len(values)
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	total := 0
	for _, v := range values {
		total += v
	}
	var dfs func(int, int) int
	dfs = func(node, fa int) int {
		loss := values[node]
		sum := 0
		for _, edge := range g[node] {
			if edge == fa {
				continue
			}
			sum += dfs(edge, node)
		}
		if sum == 0{
			return loss
		}
		return min(loss, sum)
	}
	r := dfs(0, -1)
	ans = int64(total - r)
	return
}
