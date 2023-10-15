package main

// 思路：我们可以在枚举 j 的同时，维护 nums[i] 的最大值 mx 和最小值 mn
// 在枚举j 的过程中, i = j - indexDifference也会增大，需要满足 i <= j - indexDifference (0,1,...n-1-indexDifference)
// 在遍历的过程中维护一个最大值 和 最小值即可（下标），通过下标我们能直接知道是哪个数字

func findIndices1(nums []int, indexDifference int, valueDifference int) []int {
	maxIdx, minIdx := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		i := j - indexDifference
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		} else if nums[i] < nums[minIdx] { // 这里改成了 else if， 因为一个数字不可能同时比最大值大和比最小值小
			minIdx = i
		}

		// 这里可以去掉 abs，因为如果这里的 nums[j] > 前面维护的最大值， 那么会满足下面这个 if 的条件
		// 反之亦然
		if nums[maxIdx]-nums[j] >= valueDifference {
			return []int{maxIdx, j}
		}
		if nums[j]-nums[minIdx] >= valueDifference {
			return []int{minIdx, j}
		}
	}
	return []int{-1, -1}
}

//func main() {
//	// [7],0,0
//	fmt.Println(findIndices2([]int{7}, 0, 0))
//
//	// [5,1,0,7,5,2,0], 4, 3
//	fmt.Println(findIndices2([]int{5, 1, 0, 7, 5, 2, 0}, 4, 3))
//}
