package dp

import "math"

// 1312
func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i+1][j]), float64(dp[i][j-1])) + 1)
			}
		}
	}
	return dp[0][len(s)-1]
}
