package _6__合并区间_

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	st, ed := math.MinInt32, math.MinInt32 // 哨兵（前）
	var ans [][]int

	// 添加一段哨兵（后
	intervals = append(intervals, []int{math.MaxInt32, math.MaxInt32})
	for _, interval := range intervals {
		if ed < interval[0] {
			if st != math.MinInt32 { // 如果不是 哨兵，那么将前一段区间 append
				ans = append(ans, []int{st, ed})
			}
			// 更新新的区间
			st, ed = interval[0], interval[1]
		} else {
			// 合并区间（更新 ed）
			ed = max(ed, interval[1])
		}
	}
	// 将最后一段区间 append
	//ans = append(ans, []int{st, ed})

	return ans
}
