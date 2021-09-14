package strings

import "math"

// 522
func findLUSlength(strs []string) int {
	res, n := -1, len(strs)
	for i := 0; i < n; i++ {
		if len(strs[i]) < res {
			continue
		}
		j := 0
		for j < n {
			// break when the string i is a sub seq of string j
			if i != j && isSubseq(strs[i], strs[j]) {
				break
			}
			j++
		}
		// i is not sub seq of any other strings, might be the one
		if j == n {
			res = int(math.Max(float64(res), float64(len(strs[i]))))
		}
	}
	return res
}

// abc abdec
func isSubseq(s1, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}
