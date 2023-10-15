package _38__除自身以外数组的乘积

// 思路1：前后缀分解（O(N)的空间复杂度
// pre[i] 表示 nums[i] 左边所有数的乘积
// suf[i] 表示 nums[i] 右边所有数的乘积
func productExceptSelf1(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	pre[0], suf[n-1] = 1, 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]     // pre[i] 表示 nums[i] 左边所有数的乘积
		suf[n-1-i] = suf[n-i] * nums[n-i] // suf[i] 表示 nums[i] 右边所有数的乘积
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = pre[i] * suf[i]
	}
	return ans
}

// 思路2：先算x 右边所有数的乘积，然后直接记录到答案中，再从左到右遍历一次，同时维护左边所有数的乘积（省略了前后缀乘积数组）
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	// nums: [1, 2, 3, 4]
	// ans: [24, 12, 4, 1]
	for i := n - 1; i >= 1; i-- {
		ans[i-1] = ans[i] * nums[i]
	}
	pre := 1
	for i, x := range nums {
		ans[i] *= pre // 需要的是不包括 x 的左边所有数的乘积，因此需要先乘
		pre *= x      // pre 表示 x（包括） 左边所有数的乘积
	}
	return ans
}
