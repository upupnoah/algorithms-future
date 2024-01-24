package alternating_subarray

// 双重循环
func alternatingSubarray1(nums []int) int {
	n := len(nums)
	res := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			len := j - i + 1
			if nums[j]-nums[i] != (len-1)%2 {
				break
			}
			res = max(res, len)
		}
	}
	return res
}

// 单层循环
func alternatingSubarray(nums []int) int {
	n := len(nums)
	res := -1
	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		len := j - i + 1
	// 		if nums[j]-nums[i] != (len-1)%2 {
	// 			break
	// 		}
	// 		res = max(res, len)
	// 	}
	// }
	idx := 0
	for i := 1; i < n; i++ {
		len := i - idx + 1
		if nums[i]-nums[idx] == (len-1)%2 {
			res = max(res, len)
		} else {
			if nums[i]-nums[i-1] == 1 {
				idx = i - 1
				res = max(res, 2)
			} else {
				idx = i
			}
		}
	}
	return res
}
