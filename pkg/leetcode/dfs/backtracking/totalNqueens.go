package backtracking

var res int

func totalNQueens(n int) int {
	b := make([]string, n)
	res = 0
	for i := 0; i < n; i++ {
		l := ""
		for j := 0; j < n; j++ {
			l = l + "."
		}
		b[i] = l
	}
	backTracktotal(0, b)
	return res
}

func backTracktotal(row int, board []string) {
	if row == len(board) {
		res++
	} else {
		l := []byte(board[row])
		for i := range board[row] {
			if checkValidtotal(row, i, board) {
				l[i] = 'Q'
				board[row] = string(l)
				backTracktotal(row+1, board)
				l[i] = '.'
				board[row] = string(l)
			}
		}
	}
}

func checkValidtotal(row, col int, board []string) bool {
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
