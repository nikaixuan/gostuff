package algorithms

import "math"

func zeroOneBag(bagWeight int, weights, value []int) int {
	n := len(weights)
	dp := make([][]int, len(weights)+1)
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= bagWeight; j++ {
			if bagWeight-weights[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-weights[i-1]]+value[i-1])))
			}
		}
	}
	return dp[n][bagWeight]
}

func bagProblem1d(bagWeight int, weights, value []int) int {
	n := len(weights)
	dp := make([]int, bagWeight+1)
	for i := 1; i <= n; i++ {
		for j := bagWeight; j >= 0; j-- {
			if bagWeight-weights[i-1] < 0 {
				dp[j] = dp[j]
			} else {
				dp[j] = int(math.Max(float64(dp[j]), float64(dp[j-weights[i-1]]+value[i-1])))
			}
		}
	}
	return dp[bagWeight]
}
