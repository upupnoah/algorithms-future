package number_of_boomerangs

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	cnt := map[int]int{}
	// 所有点互不相同, j = i 时, dist 为 0
	for _, p1 := range points {
		clear(cnt)
		for _, p2 := range points {
			d2 := dist(p1, p2)
			ans += cnt[d2] * 2
			cnt[d2]++
		}
	}
	return ans
}

func dist(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
