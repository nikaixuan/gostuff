package dp

// 518
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				// 用前i-1个coin凑出的方法数（已有的方法数） + 用前i个凑出 j-当前面值的方法数
				// 需要更新过的dp[j], 所以不需要改变遍历方式
				// dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]， 并非dp[i-1][j-coins[i-1]]
				// 本质：因为允许重复，所以凑（j-当前面值）不需要用前i-1个，可以包含当前面值
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[amount]
}
