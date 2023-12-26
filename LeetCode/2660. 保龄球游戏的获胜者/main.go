package is_winner

func isWinner(player1 []int, player2 []int) int {
	s1 := 0
	s2 := 0
	n := len(player1)
	for i := 0; i < n; i++ {
		if i > 0 && player1[i-1] == 10 || i > 1 && player1[i-2] == 10 {
			s1 += 2 * player1[i]
		} else {
			s1 += player1[i]
		}
		if i > 0 && player2[i-1] == 10 || i > 1 && player2[i-2] == 10 {
			s2 += 2 * player2[i]
		} else {
			s2 += player2[i]
		}
	}
	if s1 > s2 {
		return 1
	} else if s1 == s2 {
		return 0
	}
	return 2
}