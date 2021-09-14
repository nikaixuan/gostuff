package backtracking

// 216
// backtracking
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backtrackCombineSum3(k, n, []int{}, &res, 1)
	return res
}

func backtrackCombineSum3(k, n int, r []int, res *[][]int, index int) {
	if k == 0 && n == 0 {
		temp := make([]int, 0)
		*res = append(*res, append(temp, r...))
		return
	}
	for i := index; i <= 9; i++ {
		if k-1<0 || n-i<0 {
			continue
		}
		r = append(r, i)
		backtrackCombineSum3(k-1, n-i, r, res, i+1)
		r = r[:len(r)-1]
	}
}
