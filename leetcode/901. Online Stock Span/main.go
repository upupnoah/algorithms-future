package _01__Online_Stock_Span

// 算法一： 暴力

//type StockSpanner struct {
//	data []int
//}
//
//func Constructor() StockSpanner {
//	return StockSpanner{}
//}
//
//func (s *StockSpanner) Next(price int) int {
//	n := len(s.data)
//	span := 1
//	for i := n - 1; i >= 0; i-- {
//		if s.data[i] > price {
//			break
//		}
//		span++
//	}
//	s.data = append(s.data, price)
//	return span
//}

// 算法二： 单调栈

//type StockSpanner struct {
//	monoStack []int
//	prices    []int
//}
//
//func Constructor() StockSpanner {
//	return StockSpanner{}
//}

//func (s *StockSpanner) Next(price int) int {
//	s.prices = append(s.prices, price)
//	span := 1
//	if len(s.monoStack) == 0 || s.prices[s.monoStack[len(s.monoStack)-1]] > price {
//		s.monoStack = append(s.monoStack, len(s.prices)-1)
//		return span
//	}
//	for len(s.monoStack) > 0 && s.prices[s.monoStack[len(s.monoStack)-1]] <= price {
//		s.monoStack = s.monoStack[:len(s.monoStack)-1]
//	}
//	if len(s.monoStack) == 0 {
//		span = len(s.prices)
//	} else {
//		span = len(s.prices) - 1 - s.monoStack[len(s.monoStack)-1]
//	}
//	s.monoStack = append(s.monoStack, len(s.prices)-1)
//	return span
//}

// 单调栈：优雅写法

type pair struct {
	price, cnt int
}

type StockSpanner struct {
	stk []pair
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (s *StockSpanner) Next(price int) int {
	cnt := 1
	for len(s.stk) > 0 && s.stk[len(s.stk)-1].price <= price {
		cnt += s.stk[len(s.stk)-1].cnt // 当前跨度 + 栈顶元素跨度
		s.stk = s.stk[:len(s.stk)-1]
	}
	s.stk = append(s.stk, pair{price, cnt})
	return cnt
}
