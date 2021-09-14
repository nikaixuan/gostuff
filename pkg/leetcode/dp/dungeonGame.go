package dp

import "math"

// 174
// dp from array[m][n]
func calculateMinimumHP(dungeon [][]int) int {
	dp := make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[0]))
	}
	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
				if dungeon[i][j] >= 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 - dungeon[i][j]
				}
			} else if i == len(dungeon)-1 {
				if dp[i][j+1]-dungeon[i][j] <= 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j+1] - dungeon[i][j]
				}
			} else if j == len(dungeon[0])-1 {
				if dp[i+1][j]-dungeon[i][j] <= 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j] - dungeon[i][j]
				}
			} else {
				res := int(math.Min(float64(dp[i+1][j]-dungeon[i][j]), float64(dp[i][j+1]-dungeon[i][j])))
				if res <= 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = res
				}
			}
		}
	}
	return dp[0][0]
}
