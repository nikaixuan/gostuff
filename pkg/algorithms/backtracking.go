package algorithms

import (
	"fmt"
)

/*
	======================================================
	Eight Queen
*/

type board struct {
	b [8][8]int
}

func (b *board) print() {
	for i := range b.b {
		fmt.Println(b.b[i])
	}
	fmt.Println("=======================")
}
func (b *board) init() {
	b.b = [8][8]int{}
}

var eightBoard = new(board)

//var vBoard = [8]int{}
var solution = 0

func EightQueens() {
	backTrack(0)
	fmt.Println(solution)
}

func backTrack(row int) {
	if row == len(eightBoard.b) {
		solution++
		eightBoard.print()
	} else {
		for i := range eightBoard.b[row] {
			if checkValid(row, i) {
				eightBoard.b[row][i] = 1
				backTrack(row + 1)
				eightBoard.b[row][i] = 0
			}
		}
	}
}

func checkValid(row, col int) bool {
	for i := 0; i < row; i++ {
		if eightBoard.b[i][col] == 1 {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if eightBoard.b[i][j] == 1 {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if eightBoard.b[i][j] == 1 {
			return false
		}
	}
	return true
}
