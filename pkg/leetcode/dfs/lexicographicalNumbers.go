package dfs

// 386
func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i < 10; i++ {
		lexDfs(i, n, &res)
	}
	return res
}

func lexDfs(cur int, n int, r *[]int) {
	if cur > n {
		return
	}
	*r = append(*r, cur)
	for i := 0; i < 10; i++ {
		if 10*cur+i > n {
			return
		}
		lexDfs(10*cur+i, n, r)
	}
}
