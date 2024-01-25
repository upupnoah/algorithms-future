package sum_indices

func sumIndicesWithKSetBits(nums []int, k int) int {
	ans := 0
	for i, x := range nums {
		cnt := 0
		for i > 0 {
			i -= lowbit(i)
			cnt ++
		}
		if cnt == k {
			ans += x
		}
	}
	return ans
}
func lowbit(x int) int {
	return x & (-x)
}