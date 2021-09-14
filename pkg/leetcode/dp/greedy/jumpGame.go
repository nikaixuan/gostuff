package greedy

import "math"

// 55
func canJump(nums []int) bool {
	max := 0
	for i, v := range nums {
		if max < i {
			return false
		}
		max = int(math.Max(float64(max), float64(i+v)))
	}
	return true
}
