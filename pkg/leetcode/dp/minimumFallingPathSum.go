package dp

import "math"

// 931
func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		if i == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				dp[i][j] = matrix[i][j]
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[0] {
			// update the dp array to compare
			dp[i][j] = dp[i-1][j] + matrix[i][j]
			// exclude edge cases, not compare them
			if j-1 >= 0 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j-1]+matrix[i][j])))
			}
			if j < len(matrix[0])-1 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j+1]+matrix[i][j])))
			}
		}
	}
	res := dp[len(matrix)-1][0]
	for j := range dp[0] {
		res = int(math.Min(float64(res), float64(dp[len(matrix)-1][j])))
	}
	return res
}

// see the full solution, which grid related to target dp grid
func minFallingPathSumZip(matrix [][]int) int {
	dp := make([][]int, 2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
		if i == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				dp[i][j] = matrix[i][j]
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[0] {
			dp[1][j] = dp[0][j] + matrix[i][j]
			if j-1 >= 0 {
				dp[1][j] = int(math.Min(float64(dp[1][j]), float64(dp[0][j-1]+matrix[i][j])))
			}
			if j < len(matrix[0])-1 {
				dp[1][j] = int(math.Min(float64(dp[1][j]), float64(dp[0][j+1]+matrix[i][j])))
			}
		}
		for i := range dp[0] {
			dp[0][i], dp[1][i] = dp[1][i], 0
		}
	}
	res := dp[0][0]
	for j := range dp[0] {
		res = int(math.Min(float64(res), float64(dp[0][j])))
	}
	return res
}
