package dp

import "math"

// 42
func trap(height []int) int {
	a := 0
	b := len(height) - 1
	rm, lm := 0, 0
	res := 0
	for a <= b {
		lm = int(math.Max(float64(lm), float64(height[a])))
		rm = int(math.Max(float64(rm), float64(height[b])))
		if lm > rm {
			res = res + (rm - height[b])
			b--
		} else {
			res = res + (lm - height[a])
			a++
		}
	}
	return res
}
