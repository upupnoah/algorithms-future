package find_peak_grid

func findPeakGrid(mat [][]int) []int {
	n, _ := len(mat), len(mat[0])
	l, r := 0, n-1
	for l < r {
		mid := (l+r) >> 1
		t := indexOfMax(mat[mid])
		if mat[mid][t] > mat[mid+1][t] {
			r = mid
		} else {
			l = mid+1
		}
	}
	return []int{r, indexOfMax(mat[r])}
}

func indexOfMax(num []int) int {
	ans := 0
	for i, v := range num {
		if num[ans] < v {
			ans = i
		}
	}
	return ans
}