package dp

import "math"

// 221
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	result := 0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j-1]), float64(dp[i-1][j])), float64(dp[i][j-1]))) + 1
				result = int(math.Max(float64(result), float64(dp[i][j])))
			}
		}
	}
	return result * result
}
