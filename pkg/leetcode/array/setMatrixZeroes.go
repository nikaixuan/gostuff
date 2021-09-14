package array

// 73
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	fr, fc := false, false
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					fr = true
				}
				if j == 0 {
					fc = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if fr {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
	if fc {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}
