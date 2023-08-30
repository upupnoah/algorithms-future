package _654__到家的最少跳跃次数

// 思路：BFS
// 难点在于：是一个无限图，需要证明上限是多少
// max(max(forbidden)+a,x)+b
// 分情况讨论：a = b, a > b, a < b
// a < b 比较难

func minimumJumps(forbidden []int, a int, b int, x int) int {
	visited := make(map[int]bool)
	maxVal := 0
	for _, v := range forbidden {
		maxVal = max(maxVal, v)
		visited[v] = true // 不能走的点相当于已经走过了
	}
	upper := max(maxVal+a, x) + b
	type pair struct {
		v      int
		isBack bool
	}
	var q []pair
	q = append(q, pair{0, false})
	ans := 0
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			//t, isBack := q[0].v, q[0].isBack
			//q = q[1:]
			t, isBack := q[i].v, q[i].isBack

			if t == x {
				return ans
			}
			if !isBack && t-b >= 0 && !visited[t-b] {
				visited[t-b] = true
				q = append(q, pair{t - b, true})
			}
			if t+a <= upper && !visited[t+a] {
				visited[t+a] = true
				q = append(q, pair{t + a, false})
			}
		}
		q = q[levelSize:]
		ans++
	}
	return -1
}
