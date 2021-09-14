package dp

// 63
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// 滚动数组
func uniquePathsWithObstaclesRollingArray(obstacleGrid [][]int) int {
	dp := make([]int, len(obstacleGrid[0]))
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j > 0 && obstacleGrid[i][j-1] != 1 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[len(obstacleGrid[0])-1]
}
