package dp

import "math"

// 264
func nthUglyNumber(n int) int {
	u := make([]int, n)
	u[0] = 1
	f2, f3, f5 := 2, 3, 5
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		min := int(math.Min(float64(f2), math.Min(float64(f3), float64(f5))))
		u[i] = min
		if f2 == min {
			i2++
			f2 = 2 * u[i2]
		}
		if f3 == min {
			i3++
			f3 = 3 * u[i3]
		}
		if f5 == min {
			i5++
			f5 = 5 * u[i5]
		}
	}
	return u[n-1]
}
