package backtracking

func solveSudoku(board [][]byte) {
	sudokuBacktrack(board, 0, 0)
}

func sudokuBacktrack(board [][]byte, i, j int) bool {
	if j == 9 {
		return sudokuBacktrack(board, i+1, 0)
	}
	if i == 9 {
		return true
	}
	if board[i][j] != '.' {
		return sudokuBacktrack(board, i, j+1)
	}
	for n := '1'; n <= '9'; n++ {
		// 用continue可以，把条件反转，将后面的逻辑放进这个block也可以
		// 目的就是过滤不valid的答案
		if !isValidSudoku(board, i, j, byte(n)) {
			continue
		}
		board[i][j] = byte(n)
		// 回溯方法必须返回布尔，因为此题要直接修改输入的board，而不是添加结果到数组，所以一旦出现结果，必须要通过连续返回true
		// 跳出循环
		// 如果没有返回的bool和此判断block，在发现正确结果后，程序仍然会返回此层并撤销选择继续递归
		// 最后返回无解。所以如果没有此判断，必须在发现解答的时候(i==9)记录此时的board
		if sudokuBacktrack(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValidSudoku(board [][]byte, r, c int, v byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][c] == v {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[r][j] == v {
			return false
		}
	}
	for i := (r / 3) * 3; i < (r/3+1)*3; i++ {
		for j := (c / 3) * 3; j < (c/3+1)*3; j++ {
			if board[i][j] == v {
				return false
			}
		}
	}
	return true
}
