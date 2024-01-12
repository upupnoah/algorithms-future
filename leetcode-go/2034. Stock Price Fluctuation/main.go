package _034__Stock_Price_Fluctuation

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

type StockPrice struct {
	prices       *redblacktree.Tree
	timePriceMap map[int]int
	maxTimestamp int
}

func Constructor() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0}
}

func (sp *StockPrice) Update(timestamp int, price int) {
	// 由于可能有重复的股票价格，对于不支持多重有序集合（如 C++ 中的 multiset 的语言，
	// 可以额外记录每个股票价格的出现次数，在加入、删除股票价格时，更新有序集合中该股票价格的出现次数
	if prevPrice := sp.timePriceMap[timestamp]; prevPrice > 0 { // 之前记录过
		if times, _ := sp.prices.Get(prevPrice); times.(int) > 1 {
			sp.prices.Put(prevPrice, times.(int)-1)
		} else {
			sp.prices.Remove(prevPrice)
		}
	}
	times := 0
	if val, ok := sp.prices.Get(price); ok {
		times = val.(int)
	}
	sp.prices.Put(price, times+1)
	sp.timePriceMap[timestamp] = price
	sp.maxTimestamp = max(sp.maxTimestamp, timestamp)
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	return sp.prices.Right().Key.(int)
}

func (sp *StockPrice) Minimum() int {
	return sp.prices.Left().Key.(int)
}
