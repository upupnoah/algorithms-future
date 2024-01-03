package main

// 暴力
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if abs(i-j) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//func main() {
//	//nums = [5,1,4,1], indexDifference = 2, valueDifference = 4
//	nums := []int{0}
//	indexDifference := 0
//	valueDifference := 0
//	fmt.Println(findIndices(nums, indexDifference, valueDifference))
//}
