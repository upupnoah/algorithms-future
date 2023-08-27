package _7__插入区间

// 思路：
// 分成三部分看
// 1. 前面部分不相交的
// 2. 中间需要合并的
// 3. 后面部分不想交的
func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	k := 0
	n := len(intervals)
	for k < n && intervals[k][1] < newInterval[0] {
		ans = append(ans, intervals[k])
		k++
	}
	if k < n {
		newInterval[0] = min(newInterval[0], intervals[k][0])
		for k < n && newInterval[1] >= intervals[k][0] {
			newInterval[1] = max(newInterval[1], intervals[k][1])
			k++
		}
	}
	ans = append(ans, newInterval)
	for k < n {
		ans = append(ans, intervals[k])
		k++
	}
	return ans
}
