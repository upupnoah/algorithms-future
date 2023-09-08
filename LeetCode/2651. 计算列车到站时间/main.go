package _651__计算列车到站时间

func findDelayedArrivalTime(arrivalTime int, delayedtime int) int {
	return (arrivalTime + delayedtime) % 24
}
