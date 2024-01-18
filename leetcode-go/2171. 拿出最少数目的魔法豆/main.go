package minimum_removal

import "slices"

func minimumRemoval(beans []int) int64 {
	slices.Sort(beans)
	sum, mx := 0, 0
	for i, v := range beans {
		sum += v
		mx = max(mx, (len(beans)-i)*v)
	}
	return int64(sum - mx)
}
