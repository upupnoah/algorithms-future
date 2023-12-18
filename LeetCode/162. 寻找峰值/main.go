package find_peak_element


// 1. 一定有解
// 2. 二分不会错过峰值
func findPeakElement(nums []int) int {
	n := len(nums)
	// ans := sort.Search(n-1, func(x int) bool {
	// 	if nums[x] > nums[x+1] {
	// 		return true
	// 	}
	// 	return false
	// })

	// version 1
	// l, r := 0, n-1
	// for l < r {
	// 	mid := (l + r) >> 1
	// 	if nums[mid] > nums[mid+1] {
	// 		r = mid
	// 	} else {
	// 		l = mid + 1
	// 	}
	// }

	// version 2
	l, r := 0, n-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] > nums[mid-1] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
