package dp

import "math"

// 313
func nthSuperUglyNumber(n int, primes []int) int {
	u := make([]int, n)
	u[0] = 1
	f := make([]int, len(primes))
	copy(f, primes)
	index := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := f[0]
		for j := 0; j < len(primes); j++ {
			min = int(math.Min(float64(min), float64(f[j])))
		}
		u[i] = min
		for j := 0; j < len(primes); j++ {
			if min == f[j] {
				index[j]++
				f[j] = primes[j] * u[index[j]]
			}
		}
	}
	return u[n-1]
}
