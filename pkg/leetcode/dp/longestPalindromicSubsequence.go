package dp

import "math"

// 516
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[j][i] = dp[j+1][i-1] + 2
			} else {
				dp[j][i] = int(math.Max(float64(dp[j][i-1]), float64(dp[j+1][i])))
			}
		}
	}
	return dp[0][len(s)-1]
}

func longestPalindromeSubseqZip(s string) int {
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < len(s); j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre + 2
			} else {
				dp[j] = int(math.Max(float64(dp[j-1]), float64(dp[j])))
			}
			pre = temp
		}
	}
	return dp[len(s)-1]
}
