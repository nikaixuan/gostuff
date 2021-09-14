package dp

import "math"

// 152
func maxProductSubArray(nums []int) int {
	res := nums[0]
	l := 0
	r := 0
	for i := 0; i < len(nums); i++ {
		if l == 0 {
			l = nums[i]
		} else {
			l = l * nums[i]
		}
		if r == 0 {
			r = nums[len(nums)-1-i]
		} else {
			r = r * nums[len(nums)-1-i]
		}
		res = int(math.Max(float64(res), math.Max(float64(l), float64(r))))
	}
	return res
}
