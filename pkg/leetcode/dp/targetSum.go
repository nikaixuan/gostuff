package dp

// 494
// dfs(backtracking solution)
func findTargetSumWays(nums []int, target int) int {
	result := new(int)
	if len(nums) == 0 {
		return 0
	}
	sumDfs(target, nums, 0, 0, result)
	return *result
}

func sumDfs(t int, nums []int, res int, n int, result *int) {
	if n == len(nums) {
		if res == t {
			*result++
		}
		return
	}
	sumDfs(t, nums, res+nums[n], n+1, result)
	sumDfs(t, nums, res-nums[n], n+1, result)
}

// dp solution
func findTargetSumWaysDp(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < target || (sum+target)%2 == 1 {
		return 0
	}
	if target < 0 && target < 0-sum {
		return 0
	}
	return bagSubsets(nums, (sum+target)/2)
}

// bag problem
func bagSubsets(nums []int, sum int) int {
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j <= sum; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]

}

// 1-d array
func bagSubsets1d(nums []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] + dp[j-nums[i]]
			} else {
				// 此处要每次遍历j都要更新dp[j]的值
				dp[j] = dp[j]
			}
		}
	}
	return dp[sum]
}
