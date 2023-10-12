package _512__Reward_Top_K_Students

import (
	"sort"
	"strings"
)

// 模拟：哈希 + 字符串分割 split + 排序
func topStudents(positive_feedback []string, negative_feedback []string,
	report []string, student_id []int, k int) []int {
	feedPointsMap := make(map[string]int)
	for _, p := range positive_feedback {
		feedPointsMap[p] = 3
	}
	for _, n := range negative_feedback {
		feedPointsMap[n] = -1
	}
	type studentPoint struct {
		id, points int
	}
	n := len(student_id)
	sp := make([]studentPoint, n)
	for i, id := range student_id {
		p := 0
		for _, c := range strings.Split(report[i], " ") {
			p += feedPointsMap[c]
		}
		sp[i] = studentPoint{id, p}
	}
	sort.Slice(sp, func(i, j int) bool {
		a, b := sp[i], sp[j]
		return a.points > b.points || a.points == b.points && a.id < b.id
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = sp[i].id
	}
	return ans
}
