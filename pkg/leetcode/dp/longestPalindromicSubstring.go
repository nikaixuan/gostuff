package dp

// 5
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i, _ := range s {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(res) {
				res = string([]rune(s)[i : j+1])
			}
		}
	}
	return res
}
