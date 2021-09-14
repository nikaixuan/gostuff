package bfs

// 417
func pacificAtlantic(heights [][]int) [][]int {
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m := len(heights)
	n := len(heights[0])
	atlantic := make([][]bool, m)
	pacific := make([][]bool, m)
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		atlantic[i] = make([]bool, n)
		pacific[i] = make([]bool, n)
	}
	pqueue := make([][]int, 0)
	aqueue := make([][]int, 0)
	for i := 0; i < m; i++ {
		pqueue = append(pqueue, []int{i, 0})
		aqueue = append(aqueue, []int{i, n - 1})
		pacific[i][0] = true
		atlantic[i][n-1] = true
	}
	for i := 0; i < n; i++ {
		pqueue = append(pqueue, []int{0, i})
		aqueue = append(aqueue, []int{m - 1, i})
		pacific[0][i] = true
		atlantic[m-1][i] = true
	}
	bfs(heights, pqueue, pacific, dir)
	bfs(heights, aqueue, atlantic, dir)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs(matrix [][]int, queue [][]int, visited [][]bool, dir [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, v := range dir {
			x := curr[0] + v[0]
			y := curr[1] + v[1]
			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] || matrix[x][y] < matrix[curr[0]][curr[1]] {
				continue
			}
			queue = append(queue, []int{x, y})
			visited[x][y] = true
		}
	}
}
