package dp

import "math"

// 467
// bad solution
func findSubstringInWraproundString(p string) int {
	dp := make([][]int, 26)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	dp[p[0]-'a'][p[0]-'a']++
	currstart := 0
	for i := 1; i < len(p); i++ {
		// record single char
		if dp[p[i]-'a'][p[i]-'a'] == 0 {
			dp[p[i]-'a'][p[i]-'a'] = 1
		}
		// when it is not continuous, skip
		if p[i]-p[i-1] != 1 && p[i-1]-p[i] != 25 {
			currstart = i
		} else {
			for j := currstart; j < i; j++ {
				// record the longest string number
				// when "abc...zab", dp['a']['b'] should be 2 because it contains "ab" and "a...zab"
				dp[p[j]-'a'][p[i]-'a'] = int(math.Max(float64((i-j)/26+1), float64(dp[p[j]-'a'][p[i]-'a'])))
			}
		}
	}
	res := 0
	for i := range dp {
		for j := range dp[i] {
			res += dp[i][j]
		}
	}
	return res
}

// 难点在于如何想到以字母为结束的字符串数量之和等于unique non-empty substrings数量
func findSubstringInWraproundString2(p string) int {
	dp := make([]int, 26)
	dp[p[0]-'a'] = 1
	// longest string end with current byte
	longest := 1
	for i := 1; i < len(p); i++ {
		// if continue, increase the longest
		if p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25 {
			longest++
		} else {
			longest = 1
		}
		dp[p[i]-'a'] = int(math.Max(float64(dp[p[i]-'a']), float64(longest)))
	}
	res := 0
	for i := range dp {
		res += dp[i]
	}
	return res
}
