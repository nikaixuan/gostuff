package dp

import "math"

// 712
// similar to 583 and 1143
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i == 0 {
			for j := 1; j <= len(s2); j++ {
				dp[i][j] = dp[i][j-1] + int(s2[j-1])
			}
		} else {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// need to be min of the sum
				// when choose from "abb"+"a" and "aca"+"c", previous status are same but result is diff
				dp[i][j] = int(math.Min(float64(dp[i][j-1]+int(s2[j-1])), float64(dp[i-1][j]+int(s1[i-1]))))
			}
		}
	}
	return dp[m][n]
}
