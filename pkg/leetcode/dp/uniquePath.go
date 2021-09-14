package dp

// 62
func uniquePaths(m int, n int) int {
	row := make([]int, m)
	dp := make([][]int, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, row)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}

// float to int
func uniquePathsPermutation(m int, n int) int {
	res := 1
	for i, j := m+n-2, m-1; j > 0; j-- {
		res = res * i / j
		i--
	}
	return res
}
