package dp

import (
	"math"
	"sort"
)

// 120
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle[len(triangle)-1]))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = int(math.Min(float64(dp[j-1]+triangle[i][j]), float64(dp[j]+triangle[i][j])))
			}
		}
	}
	sort.Ints(dp)
	return dp[0]
}
