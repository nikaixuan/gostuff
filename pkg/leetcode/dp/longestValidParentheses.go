package dp

import "math"

// 32
func longestValidParentheses(s string) int {
	max := 0
	// dp array refers to how many continuous valid () end with current index, not the max, use max to find the max
	// ()()() or (()) are two continuous valid ()
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		// exclude all scenarios that end with '(', must be zero
		if s[i] == ')' {
			// simple (+), record 2
			if s[i-1] == '(' {
				if i < 2 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else {
				// ((())), find the index before previous valid ()s, so it is i-dp[i-1]-1
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					// whether need to plus the number of valid before this nested valid ()
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			max = int(math.Max(float64(dp[i]), float64(max)))
		}
	}
	return max
}
