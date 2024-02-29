package main

// 状态定义: 想出形如 nums 的前 i 个数是否有效划分
// 1. 如果最后两个数相等, 去掉这两个数, 问题变成剩下 n-2 个数是否有效划分
// 2. 如果 nums 的最后三个数相等, 去掉这三个数, 剩下 n-3 个数是否有效划分
// 3. 如果 nums 的最后三个数连续递增, 去掉这三个数, 剩下 n-3 个数是否有效划分

// 我的思考:
// 1. 递归做 -> 倒着来, 从 dfs(n-1) 开始
// 2. 动态规划做 -> 正着来, 从 f[0] 开始递推, f[i+1] 表示能否有效划分 nums[0] 到 nums[i], ans 为 f[n]
func validPartition(nums []int) bool {
	// 动态规划
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	// 计算 f[i+1]
	for i, x := range nums {
		if i > 0 && f[i-1] && x == nums[i-1] ||
			i > 1 && f[i-2] && (x == nums[i-1] && x == nums[i-2] ||
				x == nums[i-1]+1 && x == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}
