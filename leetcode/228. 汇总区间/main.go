package _28__汇总区间

import "strconv"

// 双指针
func summaryRanges(nums []int) []string {
	n := len(nums)
	var ans []string
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n && nums[j] == nums[j-1]+1 {
			j++
		}
		if j == i+1 {
			ans = append(ans, strconv.Itoa(nums[i]))
		} else {
			ans = append(ans, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
		}
		i = j - 1
	}
	return ans
}
