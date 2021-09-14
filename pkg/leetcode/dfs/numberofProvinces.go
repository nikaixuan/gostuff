package dfs

// 547
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	v := make(map[int]bool, n)
	r := 0
	for i := 0; i < n; i++ {
		if !v[i] {
			proDfs(isConnected, i, v)
			r++
		}
	}
	return r
}

func proDfs(connect [][]int, n int, visit map[int]bool) {
	l := len(connect)
	for i := 0; i < l; i++ {
		if !visit[i] && connect[n][i] == 1 {
			visit[i] = true
			proDfs(connect, i, visit)
		}
	}
}
