package dp

import "math"

// 132
func minCut(s string) int {
	dp := make([]int, len(s))
	par := make([][]bool, len(s))
	for i := range par {
		par[i] = make([]bool, len(s))
	}

	// "aabbc" i move first, j follow, use par array to record whether is par
	for i := 0; i < len(s); i++ {
		dp[i] = i
		for j := 0; j <= i; j++ {
			// j+1>i-1 is when "aa" or "aba", if string is longer, find it in boolean array
			if s[i] == s[j] && (j+1 > i-1 || par[j+1][i-1]) {
				par[j][i] = true
				// from start, then do not need to cut
				if j == 0 {
					dp[i] = 0
				} else {
					// max = cut before j + 1 new cut
					dp[i] = int(math.Min(float64(dp[i]), float64(dp[j-1]+1)))
				}
			}
		}
	}
	return dp[len(s)-1]
}
