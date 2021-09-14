package array

import (
	"strconv"
)

// 36
func isValidSudoku(board [][]byte) bool {
	m := make(map[string]struct{})
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				curr := "(" + string(board[i][j]) + ")"
				_, ok1 := m[strconv.Itoa(i)+curr]
				if !ok1 {
					m[strconv.Itoa(i)+curr] = struct{}{}
				} else {
					return false
				}
				_, ok2 := m[curr+strconv.Itoa(j)]
				if !ok2 {
					m[curr+strconv.Itoa(j)] = struct{}{}
				} else {
					return false
				}
				_, ok3 := m[strconv.Itoa(i/3)+curr+strconv.Itoa(j/3)]
				if !ok3 {
					m[strconv.Itoa(i/3)+curr+strconv.Itoa(j/3)] = struct{}{}
				} else {
					return false
				}
			}
		}
	}
	return true
}
