package main

import "math"

type Graph struct {
	g [][]int
}

const inf = math.MaxInt/2

func Constructor(n int, edges [][]int) Graph {
	graph := Graph{
		g: make([][]int, n),
	}
	for i := range graph.g {
		graph.g[i] = make([]int, n)
		for j := range graph.g[i] {
			graph.g[i][j] = inf
		}
	}
	for _, edge := range edges {
		x, y, w := edge[0], edge[1], edge[2]
		graph.g[x][y] = w
	}
	return graph
}

func (this *Graph) AddEdge(edge []int) {
	x, y, w := edge[0], edge[1], edge[2]
	this.g[x][y] = w
}

func (this *Graph) ShortestPath(start int, end int) int {
	n := len(this.g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range this.g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
