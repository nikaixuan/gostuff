package dfs

// 200
func numIslands(grid [][]byte) int {
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res = res + 1
				islandDfs(i, j, grid)
			}
		}
	}
	return res
}

func islandDfs(i int, j int, grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	islandDfs(i+1, j, grid)
	islandDfs(i-1, j, grid)
	islandDfs(i, j-1, grid)
	islandDfs(i, j+1, grid)
}
