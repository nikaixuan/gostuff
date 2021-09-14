package dp

import "math"

// 396
func maxRotateFunction(nums []int) int {
	max, sum := 0, 0
	for i := range nums {
		max += nums[i] * i
		sum += nums[i]
	}
	curr := max
	for i := len(nums) - 1; i >= 1; i-- {
		// F = prev F - prevLastIndex * prevLast + (sum - prevLast)
		curr = curr - (len(nums)-1)*nums[i] + (sum - nums[i])
		max = int(math.Max(float64(curr), float64(max)))
	}
	return max
}
