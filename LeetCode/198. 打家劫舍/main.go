package main

// 当前：第 i 个房子选/不选
// 子问题：从前 i 个房子中得到的最大金额和
// 下一个子问题？分类讨论：
// 不选：从前 i-1 个房子中得到的最大金额和
// 选：从前 i-2 个房子中得到的最大金额和
// dfs(i) = max(dfs(i-1), dfs(i-2) + nums[i])

// 递归写法超时
//func rob(nums []int) int {
//	n := len(nums)
//	var dfs func(int) int
//	dfs = func(i int) int {
//		if i < 0 {
//			return 0
//		}
//		return max(dfs(i-1), dfs(i-2)+nums[i])
//	}
//	return dfs(n - 1)
//}

// 递归 + 记忆化
// 时间复杂度计算公式：单个状态的时间 * 状态总数
// 时间复杂度：O(n)
// O（1）* O(n) = O(n)
//func rob(nums []int) int {
//	n := len(nums)
//	var dfs func(int) int
//	memo := make([]int, n)
//	for i := range memo {
//		memo[i] = -1
//	}
//	dfs = func(i int) int {
//		if i < 0 {
//			return 0
//		}
//		if memo[i] != -1 {
//			return memo[i]
//		}
//		res := max(dfs(i-1), dfs(i-2)+nums[i])
//		memo[i] = res
//		return res
//	}
//	return dfs(n - 1)
//}

// 自顶向下计算 = 递归 + 记忆化
// 自底向上计算 = 递推（动态规划）
// 翻译成递推（动态规划）
//
//		dfs -> f数组
//	 递归 -> 循环
//	 递归边界 -> 数组初始值
//
// f[i-1] 和 f[i-2] 归到 f[i]
// f[i] = max(f[i-1], f[i-2] + nums[i])
// 这样在 i = 0， 1 的时候，会出现负数下标，需要特殊处理
// 都+2 即可，f[i+2] = max(f[i+1], f[i] + nums[i])
//func rob(nums []int) int {
//	n := len(nums)
//	f := make([]int, n+2)
//	for i := 0; i < n; i++ {
//		f[i+2] = max(f[i+1], f[i]+nums[i])
//	}
//	return f[n+1]
//}

// 优化空间复杂度
// 在上面的基础上，f[i+2] 每次只用到 f[i+1] 和 f[i]
// f0表示上上一个，f1 表示上一个
// newF = max(f1, f0 + nums[i])
// f0 = f1
// f1 = newF
func rob(nums []int) int {
	n := len(nums)
	f0, f1, newF := 0, 0, 0
	for i := 0; i < n; i++ {
		newF = max(f1, f0+nums[i])
		f0 = f1
		f1 = newF
	}
	return newF
}
