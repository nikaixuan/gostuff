package dp

import "math"

// 376
func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return int(math.Max(float64(up), float64(down)))
}
