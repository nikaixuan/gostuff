package dp

import (
	"math"
)

// 72
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	// init
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i == 0 {
			for j := 1; j < len(dp[0]); j++ {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := math.Min(float64(dp[i][j-1]), float64(dp[i-1][j]))
				dp[i][j] = int(math.Min(min, float64(dp[i-1][j-1])) + 1)
			}
		}
	}
	return dp[m][n]
}
