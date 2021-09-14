package dfs

// 79
func exist(board [][]byte, word string) bool {
	v := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		v[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if wordSearchDfs(board, v, []byte(word), i, j, 0) {
				return true
			}
		}
	}
	return false
}

func wordSearchDfs(board [][]byte, visited [][]bool, word []byte, i, j, n int) bool {
	if board[i][j] == word[n] {
		if len(word)-1 == n {
			return true
		}
		visited[i][j] = true
		if i > 0 && !visited[i-1][j] {
			r := wordSearchDfs(board, visited, word, i-1, j, n+1)
			if r {
				return r
			}
		}
		if i < len(board)-1 && !visited[i+1][j] {
			r := wordSearchDfs(board, visited, word, i+1, j, n+1)
			if r {
				return r
			}
		}
		if j > 0 && !visited[i][j-1] {
			r := wordSearchDfs(board, visited, word, i, j-1, n+1)
			if r {
				return r
			}
		}
		if j < len(board[0])-1 && !visited[i][j+1] {
			r := wordSearchDfs(board, visited, word, i, j+1, n+1)
			if r {
				return r
			}
		}
		visited[i][j] = false
	}
	return false
}
