package maximum_obtainable_points

// 由于是从两边开始选择 k 个数字, 使得和最大
// 等价于: 从中间选择一段连续的数字(n - k 个), 使得和最小
func maxScore(cardPoints []int, k int) int {
	// 滑动窗口
	n := len(cardPoints)
	windowSize := n - k
	sum := 0
	minSum := 0
	for i, v := range cardPoints {
		if i < windowSize {
			sum += v
			minSum = sum
		} else {
			sum += v - cardPoints[i-windowSize]
			minSum = min(minSum, sum)
		}
	}
	total := 0
	for _, v := range cardPoints {
		total += v
	}
	return total - minSum
}
