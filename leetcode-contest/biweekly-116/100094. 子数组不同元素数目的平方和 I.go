package biweekly_116

func sumCounts1(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	for i := 0; i < n; i++ {
		m := make(map[int]struct{})
		for j := i; j < n; j++ {
			m[nums[j]] = struct{}{}
			cnt := len(m)
			ans = (ans + cnt*cnt) % mod
		}
	}
	return
}
