package bst

// 96
func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[n]

}
