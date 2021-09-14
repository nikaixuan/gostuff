package dfs

// 130
func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	for i := range board {
		if board[i][0] == 'O' {
			surroundDfs(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			surroundDfs(board, i, len(board[0])-1)
		}
	}
	for i := range board[0] {
		if board[0][i] == 'O' {
			surroundDfs(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			surroundDfs(board, len(board)-1, i)
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}

func surroundDfs(b [][]byte, i int, j int) {
	if i < 0 || j < 0 || i > len(b)-1 || j > len(b[i]) || b[i][j] != 'O' {
		return
	}
	if b[i][j] == 'O' {
		b[i][j] = 'Y'
	}
	surroundDfs(b, i+1, j)
	surroundDfs(b, i, j-1)
	surroundDfs(b, i-1, j)
	surroundDfs(b, i, j+1)
}
