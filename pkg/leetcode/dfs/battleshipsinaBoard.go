package dfs

// 419
func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				res++
				battleShipDfs(board, i, j)
			}
		}
	}
	return res
}

func battleShipDfs(board [][]byte, m, n int) {
	if m < 0 || n < 0 || m >= len(board) || n >= len(board[0]) || board[m][n] == '.' {
		return
	}
	board[m][n] = '.'
	battleShipDfs(board, m+1, n)
	battleShipDfs(board, m, n+1)
}
