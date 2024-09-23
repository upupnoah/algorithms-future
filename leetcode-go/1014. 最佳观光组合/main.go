package main

func maxScoreSightseeingPair(values []int) (ans int) {
	mx := 0 // j 左边的 values[i] + i 的最大值
	for j, v := range values {
		ans = max(ans, mx+v-j)
		mx = max(mx, v+j)
	}
	return
}
