package _109__Corporate_Flight_Bookings

// 算法：差分
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, 0, 2e4+1)
	for _, b := range bookings {
		diff[b[0]] += b[2]
		diff[b[1]+1] -= b[2]
	}
	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]
	}
	return diff
}
