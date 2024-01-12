package _103__环和杆

// 暴力枚举
func countPoints(rings string) (ans int) {
	st := [10][3]int{}
	for i := 0; i < len(rings); i += 2 {
		t := 0
		if rings[i] == 'G' {
			t = 1
		} else if rings[i] == 'B' {
			t = 2
		}
		st[rings[i+1]-'0'][t]++
	}
	for _, v := range st {
		cnt := 0
		for _, v2 := range v {
			if v2 > 0 {
				cnt++
			}
		}
		if cnt == 3 {
			ans++
		}
	}
	return ans
}

// 位运算优化
func countPoints2(rings string) (ans int) {
	mapping := map[byte]int{
		'R': 1,
		'G': 2,
		'B': 4,
	}
	st := [10]int{}
	for i := 0; i < len(rings); i += 2 {
		st[rings[i+1]-'0'] |= mapping[rings[i]]
	}
	for _, v := range st {
		if v == 7 {
			ans++
		}
	}
	return
}
