package can_see_persons_count

import "math"

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	stk := []int{math.MaxInt}
	for i := n - 1; i >= 0; i-- {
		for stk[len(stk)-1] < heights[i] {
			stk = stk[:len(stk)-1]
			ans[i]++
		}
		if len(stk) > 1 { // 空为 1(哨兵), 因此有一个元素就是 > 1
			ans[i]++
		}
		stk = append(stk, heights[i])
	}
	return ans
}
