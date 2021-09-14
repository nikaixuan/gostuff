package dp

// 887
// know the basic dp and binary
func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	m := 1
	for ; dp[k][m] < n; m++ {
		for i := 0; i <= k; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}
	return m
}
