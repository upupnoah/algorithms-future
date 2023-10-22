package main

const mx = 200 // the maximum length of s
var divisors [mx + 1][]int

// Preprocessing the proper factors of 1 to mx.
func init() {
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			divisors[j] = append(divisors[j], i) // j has i as a proper factor
		}
	}
}

func calc(s string) int {
	n := len(s)
	res := n
	for _, d := range divisors[n] {
		cnt := 0
		for i0 := 0; i0 < d; i0++ { // 0 ~ d-1， 分成 d 组
			var t []byte
			for i := i0; i < n; i += d {
				t = append(t, s[i])
			}
			//for i, m := 0, len(t); i < m/2; i++ {
			//	v, w := t[i], t[m-1-i]
			//	if v != w {
			//		cnt++
			//	}
			//}
			i, j := 0, len(t)-1
			for i < j {
				if t[i] != t[j] {
					cnt++
				}
				i++
				j--
			}
		}
		res = min(res, cnt)
	}
	return res
}

func minimumChanges(s string, k int) (ans int) {
	n := len(s)
	modify := make([][]int, n-1)
	for l := range modify {
		modify[l] = make([]int, n)
		for r := l + 1; r < n; r++ {
			modify[l][r] = calc(s[l : r+1])
		}
	}
	memo := make([][]int, k)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return modify[i][j]
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := n
		defer func() { *p = res }()
		for L := i * 2; L < j; L++ {
			res = min(res, dfs(i-1, L-1)+modify[L][j])
		}
		return res
	}
	return dfs(k-1, n-1)
}
