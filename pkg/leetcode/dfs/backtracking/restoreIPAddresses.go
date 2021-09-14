package backtracking

import "strconv"

// 93
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	if len(s) > 12 {
		return res
	}
	ipBacktracking(s, &res, 4, 0, "")
	return res
}

func ipBacktracking(s string, res *[]string, k, start int, curr string) {
	if k == 0 {
		// current format should be "xx.xx.xx.xx."
		if len(curr) == len(s)+4 {
			temp := curr
			// delete the last extra '.'
			*res = append(*res, temp[:len(temp)-1])
			return
		}
	}
	// max length is 3
	for i := start + 1; i <= start+3; i++ {
		// index bound check
		if i <= len(s) {
			// if start index value is 0, that slice of ip can only contain one 0, directly go to next slice and break
			if s[start] == '0' {
				ipBacktracking(s, res, k-1, start+1, curr+string(s[start])+".")
				break
			}
			ip, _ := strconv.Atoi(s[start:i])
			if ip > 0 && ip < 256 {
				// finish one slice of ip, so k-1, the ip is from [start, i), so next start index should be i
				ipBacktracking(s, res, k-1, i, curr+strconv.Itoa(ip)+".")
			}
		}
	}
}
