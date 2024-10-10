package main

func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt := map[int]int{}
	for _, x := range nums1 {
		if x%k > 0 {
			continue
		}
		x /= k
		for d := 1; d*d <= x; d++ { // 枚举因子
			if x%d == 0 {
				cnt[d]++ // 统计因子
				if d*d < x {
					cnt[x/d]++ // 因子总是成对出现
				}
			}
		}
	}

	for _, x := range nums2 {
		ans += int64(cnt[x])
	}
	return
}
