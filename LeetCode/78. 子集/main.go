package _8__子集

// 子集型回溯
// 视角：输入的视角（选/不选）
func subsets1(nums []int) (ans [][]int) {
	n := len(nums)
	var dfs func(int)
	var path []int
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]

		// 不选
		dfs(i + 1)
	}
	dfs(0)
	return ans
}

// 答案的视角（选哪个数）
// 只有一种选择（选择他）
// 0: []
// 1: [1],[1,2],[1,2 3], [1,3]
// 2: [2],[2,3]
// 3: [3]
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...))
		//if i == n {
		//	return
		//}
		// 有可能重复，因此规定一个顺序
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0)
	return ans
}
