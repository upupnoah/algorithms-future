package main

func findChampion(n int, edges [][]int) (ans int) {
	cache := make([]int, n)
	for _, edge := range edges {
		cache[edge[1]]++
	}
	flag := false
	for i, v := range cache {
		if flag && v == 0 {
			return -1
		}
		if v == 0 {
			ans = i
			flag = true
		}

	}
	return
}
