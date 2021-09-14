package dp

import (
	"math"
)

// 787
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if j == src {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt64
			}
		}
	}
	for i := 1; i <= k+1; i++ {
		for _, flight := range flights {
			u, v, w := flight[0], flight[1], flight[2]
			if dp[i-1][u] != math.MaxInt64 {
				dp[i][v] = int(math.Min(float64(dp[i-1][v]), float64(dp[i-1][u]+w)))
			}
		}
	}
	if dp[k+1][dst] == math.MaxInt64 {
		return -1
	}
	return dp[k+1][dst]
}
