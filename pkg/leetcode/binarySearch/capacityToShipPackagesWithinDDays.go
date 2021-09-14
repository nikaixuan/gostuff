package binarySearch

import "math"

// 1011
func shipWithinDays(weights []int, days int) int {
	left := 0
	right := 0
	for _, v := range weights {
		left = int(math.Max(float64(left), float64(v)))
		right += v
	}
	for left <= right {
		mid := left + (right-left)/2
		day := getDays(weights, mid)
		if day == days {
			right = mid - 1
		} else if day > days {
			left = mid + 1
		} else if day < days {
			right = mid - 1
		}
	}
	return left
}

// the bigger the cargo, the smaller the days
func getDays(weights []int, cargo int) int {
	days := 1
	c := cargo
	for _, w := range weights {
		if c < w {
			days++
			c = cargo
		}
		c = c - w
	}
	return days
}
