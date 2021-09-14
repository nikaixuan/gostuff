package backtracking

func solveNQueens(n int) [][]string {
	b := make([]string, n)
	res := make([][]string, 0)
	for i := 0; i < n; i++ {
		l := ""
		for j := 0; j < n; j++ {
			l = l + "."
		}
		b[i] = l
	}
	backTrack(0, b, &res)
	return res
}

func backTrack(row int, board []string, r *[][]string) {
	if row == len(board) {
		newb := make([]string, len(board))
		copy(newb, board)
		*r = append(*r, newb)
	} else {
		l := []byte(board[row])
		for i := range board[row] {
			if checkValid(row, i, board) {
				l[i] = 'Q'
				board[row] = string(l)
				backTrack(row+1, board, r)
				l[i] = '.'
				board[row] = string(l)
			}
		}
	}
}

func checkValid(row, col int, board []string) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board[i]); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
