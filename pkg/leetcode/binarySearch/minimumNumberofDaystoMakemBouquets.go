package binarySearch

import (
	"math"
)

// 1482
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	min, max := math.MaxInt64, math.MinInt64
	for i := range bloomDay {
		min = int(math.Min(float64(min), float64(bloomDay[i])))
		max = int(math.Max(float64(max), float64(bloomDay[i])))
	}
	for min <= max {
		mid := min + (max-min)/2
		if checkMin(bloomDay, m, k, mid) {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return max + 1
}

func checkMin(bloomDay []int, m, k, mid int) bool {
	count := 0
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] <= mid {
			count++
		} else {
			count = 0
		}
		if count == k {
			m--
			count = 0
		}
		if m == 0 {
			return true
		}
	}
	return false
}
