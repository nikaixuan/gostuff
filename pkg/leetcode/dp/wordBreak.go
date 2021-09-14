package dp

// 139
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && stringContain(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func stringContain(s string, dic []string) bool {
	for _, v := range dic {
		if v == s {
			return true
		}
	}
	return false
}
