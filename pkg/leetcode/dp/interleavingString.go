package dp

// 97
// Not working with "aabcc", "dbbca", "aadbbcbcac"
/*
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	s1r := []rune(s1)
	s2r := []rune(s2)
	s3r := []rune(s3)
	m, n := 0, 0
	for i:=0;i< len(s3r);i++ {
		if m<len(s1r) && s3r[i] == s1r[m] {
			m = m+1
		} else if n<len(s2r) && s3r[i] == s2r[n] {
			n = n+1
		} else {
			return false
		}
	}
	return true
}
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	s1r := []rune(s1)
	s2r := []rune(s2)
	s3r := []rune(s3)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1r[i-1] == s3r[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2r[j-1] == s3r[p])
			}
		}
	}
	return dp[m][n]
}
