package dp

// 377
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	// dp[i] = dp[i-nums[0]] + dp[i-nums[1]]....
	// dp length is the target number
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		sum := 0
		for _, n := range nums {
			if i-n >= 0 {
				sum += dp[i-n]
			}
		}
		dp[i] = sum
	}
	return dp[target]
}
