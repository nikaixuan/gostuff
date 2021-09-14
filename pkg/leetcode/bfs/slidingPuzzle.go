package bfs

import (
	"strconv"
)

func slidingPuzzle(board [][]int) int {
	queue := make([]string, 0)
	queue = append(queue, getBoardString(board))
	res := -1
	m := make(map[string]struct{})
	// 当0的位置是i时，可能交换的字符串位置
	dir := [][]int{{1, 3}, {0, 2, 4},
		{1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

	for len(queue) > 0 {
		l := len(queue)
		res++
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == "123450" {
				return res
			}
			if _, ok := m[curr]; ok {
				continue
			}
			m[curr] = struct{}{}
			idx := 0
			// 找到0的位置
			for i := range curr {
				if curr[i] == '0' {
					idx = i
				}
			}
			// 和可能的位置交换
			for _, v := range dir[idx] {
				queue = append(queue, swapIdx(curr, v, idx))
			}
		}
	}
	return -1
}

func getBoardString(board [][]int) string {
	r := ""
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			r = r + strconv.Itoa(board[i][j])
		}
	}
	return r
}

func swapIdx(curr string, i, j int) string {
	temp := []byte(curr)
	temp[i], temp[j] = temp[j], temp[i]
	return string(temp)
}
