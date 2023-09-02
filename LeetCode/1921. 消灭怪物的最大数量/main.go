package _921__消灭怪物的最大数量

import (
	"sort"
)

// 思路：
// 对于每一个d[i]/s[i] 上取整，就是每一个敌人到达城市，需要的秒数
// 排序， 贪心
func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		m[i] = (dist[i] + speed[i] - 1) / speed[i]
	}
	sort.Ints(m)
	res := 1
	for res < n && res < m[res] {
		res++
	}
	return res
}
