package main

const N = 100100
const M = 200100

var h [N]int
var e [M]int
var ne [M]int
var idx int
var sum [N]int
var cnt [N]int
var up [N]int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs1(u, father int) {
	sum[u] = 0
	cnt[u] = 1

	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == father {
			continue
		}
		dfs1(j, u)
		sum[u] += sum[j] + cnt[j]
		cnt[u] += cnt[j]
	}
}

func dfs2(u, father int) {
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == father {
			continue
		}
		up[j] = up[u] + sum[u] - (sum[j] + cnt[j]) + cnt[0] - cnt[j]
		dfs2(j, u)
	}
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	idx = 0

	for _, e := range edges {
		a, b := e[0], e[1]
		add(a, b)
		add(b, a)
	}

	dfs1(0, -1)
	dfs2(0, -1)

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, sum[i]+up[i])
	}
	return res
}
