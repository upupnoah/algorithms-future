package main

func minGroupsForValidAssignment(nums []int) int {
	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums) // 最多有 n 个为（1）的组
	for _, c := range cnt {
		k = min(k, c)
	}
	for {
		ans := 0
		for _, c := range cnt {
			if c/k < c%k { // 不能分组
				ans = 0
				break
			}
			ans += (c + k) / (k + 1) // c/(k+1) 向上取整
		}
		if ans > 0 {
			return ans
		}
		k--
	}
}

// func main() {
// 	//输入：nums = [3,2,3,2,3]
// 	//输出：2
// 	fmt.Println(minGroupsForValidAssignment([]int{3, 2, 3, 2, 3}))

// 	//输入：nums = [10,10,10,3,1,1]
// 	//输出：4
// 	fmt.Println(minGroupsForValidAssignment([]int{10, 10, 10, 3, 1, 1}))

// 	//输入：nums = [1,1,3,3,1,1,2,2,3,1,3,2]
// 	//输出：5
// 	fmt.Println(minGroupsForValidAssignment([]int{1, 1, 3, 3, 1, 1, 2, 2, 3, 1, 3, 2}))
// }
