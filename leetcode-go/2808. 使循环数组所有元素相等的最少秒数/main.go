package minimum_seconds

// 1. 扩散(左边 x, 右边 x), 分别往左右扩散
// 2. 最后一定是 初始 nums 中的某个数
// 3. 处理循环数组可以 破环成链, 直接复制一份, 两个数组拼接
func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i , v := range nums {
		pos[v] = append(pos[v], i)
	}
	n := len(nums)
	ans := n // 最多 n 秒
	for _, p := range pos {
		p = append(p, p[0] + n) // 处理循环数组
		mx := 0
		for i := 1; i < len(p); i++ {
			mx = max(mx, (p[i] - p[i-1])/2)
		}
		ans = min(ans, mx)
	}
	return ans
}
