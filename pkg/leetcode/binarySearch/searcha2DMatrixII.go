package binarySearch

// 240
func searchMatrix(matrix [][]int, target int) bool {
	col, row := len(matrix[0])-1, 0
	for col >= 0 && row <= len(matrix)-1 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}
