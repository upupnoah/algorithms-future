package main

func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			ans += min(tickets[i], tickets[k])
		} else {
			ans += min(tickets[i], tickets[k]-1)
		}
	}
	return ans
}
