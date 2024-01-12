package _605__从两个数字数组里生成最小数字

func minNumber(nums1 []int, nums2 []int) int {
	m1, m2 := 10, 10
	st := make([]int, 10)
	for _, v := range nums1 {
		m1 = min(m1, v)
		st[v]++
	}
	minSame := 10
	var flag bool
	for _, v := range nums2 {
		if st[v] > 0 {
			flag = true
			minSame = min(minSame, v)
		}
		m2 = min(m2, v)
	}
	if flag {
		return minSame
	}
	if m1 > m2 {
		m1, m2 = m2, m1
	}
	return m1*10 + m2
}
