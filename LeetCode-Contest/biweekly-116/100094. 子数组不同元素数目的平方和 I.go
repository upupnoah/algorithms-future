package biweekly_116

func sumCounts1(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	//s := make([]int, n+1)
	//m := make(map[int]bool)
	//for i := 0; i < n; i++ {
	//	if ok, _ := m[nums[i]]; ok {
	//		s[i+1] = s[i]
	//	} else {
	//		s[i+1] = s[i] + 1
	//		m[nums[i]] = true
	//	}
	//}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			m := make(map[int]bool)
			cnt := 0
			for k := i; k <= j; k++ {
				if ok, _ := m[nums[k]]; !ok {
					cnt++
					m[nums[k]] = true
				}
			}
			ans = (ans%mod + (cnt*cnt)%mod) % mod
		}
	}
	return
}
