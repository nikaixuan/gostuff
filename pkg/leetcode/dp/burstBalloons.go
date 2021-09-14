package dp

import "math"

// 312
func maxCoins(nums []int) int {
	dp := make([][]int, len(nums)+2)
	for i := range dp {
		dp[i] = make([]int, len(nums)+2)
	}
	newn := make([]int, 0)
	newn = append(newn, 1)
	newn = append(newn, nums...)
	newn = append(newn, 1)
	for i := len(nums); i >= 0; i-- {
		for j := i + 1; j < len(nums)+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = int(math.Max(float64(dp[i][j]), float64(dp[i][k]+dp[k][j]+newn[i]*newn[j]*newn[k])))
			}
		}
	}
	return dp[0][len(nums)+2]
}
