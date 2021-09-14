package dp

import "math"

// 343
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = i - 1
		for j := 1; j < i; j++ {
			dp[i] = int(math.Max(float64(dp[i]), math.Max(float64(dp[j]), float64(j))*math.Max(float64(dp[i-j]), float64(i-j))))
		}
	}
	return dp[n]
}
