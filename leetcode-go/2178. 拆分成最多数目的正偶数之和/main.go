package main

// 贪心
func maximumEvenSplit(finalSum int64) []int64 {
	var res []int64
	if finalSum%2 ==1 {
		return res
	}
	for i := int64(2); i <= finalSum; i+=2 {
		res = append(res, i)
		finalSum -= i
	}
	res[len(res)-1] += finalSum
	return res
}