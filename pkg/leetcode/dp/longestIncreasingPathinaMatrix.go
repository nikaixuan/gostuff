package dp

import "math"

// 329
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	max := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// dfs every node
			l := dfsPath(matrix, &dp, i, j, m, n)
			max = int(math.Max(float64(max), float64(l)))
		}
	}
	return max
}

func dfsPath(matrix [][]int, dp *[][]int, a, b, m, n int) int {
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	// important!! exclude duplicate
	if (*dp)[a][b] != 0 {
		return (*dp)[a][b]
	}
	max := 1
	for _, dir := range dirs {
		x := a + dir[0]
		y := b + dir[1]
		if (x < m && x >= 0 && y < n && y >= 0) && matrix[x][y] > matrix[a][b] {
			l := 1 + dfsPath(matrix, dp, x, y, m, n)
			max = int(math.Max(float64(l), float64(max)))
		}
	}
	(*dp)[a][b] = max
	return max
}
