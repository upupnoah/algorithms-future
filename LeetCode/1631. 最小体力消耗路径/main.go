package minimum_effort_path

import "sort"

// 二分 + BFS
// 将问题转化为判定性问题: 是否存在一条从左上角到右下角的路径，其经过的所有边权的最大值不超过 x ？
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])

	// 上右下左

	return sort.Search(1e6, func(maxHeightDiff int) bool {
		vis := make([][]bool, n) // 表示是否已经访问过
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[0][0] = true
		queue := []pair{{0, 0}}
		for len(queue) > 0 {
			// 取出队头
			p := queue[0]
			queue = queue[1:]
			if p.x == n-1 && p.y == m-1 {
				return true
			}
			for _, d := range directions {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
					vis[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

type pair struct {
	x, y int
}

// var dx = []int{0, 1, 0, -1}
// var dy = []int{1, 0, -1, 0}
var directions = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
