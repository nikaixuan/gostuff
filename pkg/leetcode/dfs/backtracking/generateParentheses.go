package backtracking

// 22
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	parenthesisBacktrack(n, n, &res, "")
	return res
}

func parenthesisBacktrack(open, close int, res *[]string, r string) {
	if close < open {
		return
	}
	if open < 0 || close < 0 {
		return
	}
	if open == 0 && close == 0 {
		temp := r
		*res = append(*res, temp)
		return
	}
	// 也可以用for循环并把括号字符合法性检测拿出来
	r = r + "("
	parenthesisBacktrack(open-1, close, res, r)
	r = r[:len(r)-1]

	r = r + ")"
	parenthesisBacktrack(open, close-1, res, r)
	r = r[:len(r)-1]
}
