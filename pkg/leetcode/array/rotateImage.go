package array

// 48
func rotate(matrix [][]int) {
	a := 0
	b := len(matrix) - 1
	for a < b {
		tmp := matrix[b]
		matrix[b] = matrix[a]
		matrix[a] = tmp
		a++
		b--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}
