package dp

// 413
// 能化解成最优子问题，a[i:j]是ari的话，a[j+1]-a[j]将决定a[i:j+1]是不是ari
// 但还可以简化，dp窗口限制在3个元素就可以
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := make([]int, len(nums)-2)
	isAri := make([]int, len(nums)-2)

	if isArithmetic(nums[0:3]) {
		dp[0] = 1
		isAri[0] = 1
	} else {
		dp[0] = 0
	}
	for i := 1; i < len(nums)-2; i++ {
		if isArithmetic(nums[i : i+3]) {
			isAri[i] = isAri[i-1] + 1
			dp[i] = dp[i-1] + isAri[i]
		} else {
			isAri[i] = 0
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

func numberOfArithmeticSlicesLessSpace(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp1, dp2 := 0, 0
	isAri1, isAri2 := 0, 0

	if isArithmetic(nums[0:3]) {
		dp2, dp1 = 1, 1
		isAri1 = 1
	}
	for i := 1; i < len(nums)-2; i++ {
		if isArithmetic(nums[i : i+3]) {
			isAri2 = isAri1 + 1
			dp2 = dp1 + isAri2
			isAri1, dp1 = isAri2, dp2
		} else {
			isAri1 = 0
		}
	}
	return dp2
}

func isArithmetic(nums []int) bool {
	return nums[2]-nums[1] == nums[1]-nums[0]
}
