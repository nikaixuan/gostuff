package backtracking

// 17
var list = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var result []string
	var slice []int // empty slice
	for _, digit := range digits {
		slice = append(slice, int(digit)-int('0')) // build up slice
	}
	if len(slice) == 0 {
		return result
	}
	var res = ""
	backTracking(0, res, &result, slice)
	return result
}

func backTracking(n int, r string, res *[]string, s []int) {
	if n == len(s) {
		*res = append(*res, r)
	} else {
		for j := range list[s[n]-2] {
			newString := r + string(list[s[n]-2][j])
			backTracking(n+1, newString, res, s)
		}
	}
}
