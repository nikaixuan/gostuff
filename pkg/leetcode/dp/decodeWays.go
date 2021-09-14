package dp

// 91
func numDecodings(s string) int {
	pre := 1
	curr := 1
	if s[0] == '0' {
		return 0
	}
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
				curr = tmp + pre
			} else {
				curr = tmp
			}
		}
		pre = tmp
	}
	return curr
}
