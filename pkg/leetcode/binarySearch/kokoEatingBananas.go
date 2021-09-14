package binarySearch

// 875
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1000000000 + 1
	for left < right {
		mid := left + (right-left)/2
		hour := getHour(piles, mid)
		if hour <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// the bigger the x is, the small the hour is
func getHour(piles []int, x int) int {
	hours := 0
	for _, p := range piles {
		hours += p / x
		if p%x > 0 {
			hours++
		}
	}
	return hours
}
