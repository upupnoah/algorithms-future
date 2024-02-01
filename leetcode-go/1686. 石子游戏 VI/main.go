package stone_game_vi

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	sum := make([]int, n)
	for i := range aliceValues {
		sum[i] = aliceValues[i] + bobValues[i]
	}
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return sum[idx[i]] > sum[idx[j]]
	})
	a, b := 0, 0
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			a += aliceValues[idx[i]]
		} else {
			b += bobValues[idx[i]]
		}
	}
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
