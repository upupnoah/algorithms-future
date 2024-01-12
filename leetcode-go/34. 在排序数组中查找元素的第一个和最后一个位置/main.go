package _4__在排序数组中查找元素的第一个和最后一个位置

// st 0 表示找第一个， st 1 表示找最后一个
func lower_bound(nums []int, target int, st int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	if st == 0 {
		for left < right {
			mid := left + (right-left)>>1 // 避免溢出
			if nums[mid] >= target {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}
	if st == 1 {
		for left < right {
			mid := left + (right-left+1)>>1
			if nums[mid] <= target {
				left = mid
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}

func searchRange(nums []int, target int) []int {
	return []int{lower_bound(nums, target, 0), lower_bound(nums, target, 1)}
}
