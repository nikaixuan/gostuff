package backtracking

// 131
func partition(s string) [][]string {
	if len(s) == 0 {
		return make([][]string, 0)
	}
	res := make([][]string, 0)
	partitionBackTrack(s, []string{}, &res, 0)
	return res
}

func partitionBackTrack(s string, step []string, res *[][]string, start int) {
	if len(s) == start {
		temp := append([]string{}, step...)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if !isPalindromeIndex(start, i, s) {
			continue
		}
		temp := s[start:i+1]
		step = append(step, temp)
		partitionBackTrack(s, step, res, i+1)
		step = step[:len(step)-1]
	}

	// 另一种方法，先拿出字符串再检查是否回文，所以i要是start+1，且最后会成为越界索引，因为切片是开区间
	// 所以递归方法中的n直接是i，不用+1
	//for i := start+1; i <= len(s); i++ {
	//	temp := s[start:i]
	//	if !isPalindrome(temp) {
	//		continue
	//	}
	//	step = append(step, temp)
	//	partitionBackTrack(s, step, res, i)
	//	step = step[:len(step)-1]
	//}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindromeIndex(i, j int, s string) bool {
	for i<=j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
