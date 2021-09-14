package binarySearch

import "math"

// 410
func splitArray(nums []int, m int) int {
	left, right := 0, 0
	for _, v := range nums {
		// max individual
		left = int(math.Max(float64(left), float64(v)))
		// sum of all number
		right = right + v
	}
	for left <= right {
		// possible sum, need to minimize it
		mid := left + (right-left)/2
		max := minCountSplit(nums, mid)
		if max <= m {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// the bigger the max is, the less subarray the array can split
// because it needs more number to sum to max
func minCountSplit(nums []int, max int) int {
	m := 0
	count := 1
	for _, v := range nums {
		if m+v > max {
			count++
			m = v
		} else {
			m += v
		}
	}
	return count
}
