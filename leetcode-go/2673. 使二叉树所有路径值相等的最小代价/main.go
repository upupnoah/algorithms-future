package main

func minIncrements(n int, cost []int) (ans int) {
	// start counting from the last non-leaf node
	for i := n / 2; i > 0; i-- {
		left, right := cost[i*2-1], cost[i*2]
		if left > right {
			left, right = right, left
		}
		ans += right - left // change different node to same
		cost[i-1] += right  // cumulative path sum
	}
	return

}
