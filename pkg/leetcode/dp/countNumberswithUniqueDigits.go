package dp

// 357
func countNumbersWithUniqueDigits(n int) int {
	dp0, dp1 := 1, 1

	for i := 1; i <= n; i++ {
		n := 9
		for j := 1; j < i; j++ {
			n = n * (10 - j)
		}
		dp1 = dp0 + n
		dp0 = dp1
	}
	return dp1
}
