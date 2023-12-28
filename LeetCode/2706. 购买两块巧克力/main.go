package buy_choco

func buyChoco(prices []int, money int) int {
	s1, s2 := 200, 200
	for _, v := range prices {
		if v < s1 {
			s2 = s1
			s1 = v
		} else if v < s2 {
			s2 = v
		}
	}
	if s1+s2 > money {
		return money
	} else {
		return money - s1 - s2
	}
}
