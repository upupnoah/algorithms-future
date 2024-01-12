package repeat_limited_string

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}
	ans := []byte{}
	for i := 25; i >= 0; i-- {
		k := i - 1
		for {
			for j := 0; j < repeatLimit && cnt[i] > 0; j++ {
				cnt[i]--
				ans = append(ans, 'a'+byte(i))
			}
			if cnt[i] == 0 {
				break
			}
			for k >= 0 && cnt[k] == 0 {
				k--
			}
			if k < 0 {
				break
			}

			// i 没用完, 插入一个字母 k,  这样可以继续填 i
			cnt[k]--
			ans = append(ans, 'a'+byte(k))
		}
	}
	return string(ans)
}
