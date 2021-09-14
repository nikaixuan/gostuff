package dp

// 416
// sub-set bag problem
// find the bag problem in the subset of the given array
// 要转化为背包问题
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	weight := sum / 2
	dp := make([]bool, weight+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := weight; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j-nums[i-1]] || dp[j]
			}
		}
	}
	return dp[weight]
}
