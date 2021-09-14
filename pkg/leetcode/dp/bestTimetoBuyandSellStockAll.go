package dp

import "math"

// 121
func maxProfitOneTime(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(-prices[i])))
	}
	return dp[len(dp)-1][0]
}

// 122
func maxProfitInfTimes(prices []int) int {
	dp0, dp1 := 0, math.MinInt64

	for i := 0; i < len(prices); i++ {
		temp := dp0
		dp0 = int(math.Max(float64(dp0), float64(dp1+prices[i])))
		dp1 = int(math.Max(float64(dp1), float64(temp-prices[i])))
	}
	return dp0
}

// 123
func maxProfitTrade2times(prices []int) int {

	// 3-d dp array

	//dp := make([][][]int, len(prices))
	//for i := range dp {
	//	dp[i] = make([][]int, 3)
	//	for j:= range dp[i] {
	//		dp[i][j] = make([]int, 2)
	//	}
	//}
	//for i:=0;i<len(prices);i++ {
	//	for k:=1;k<=2;k++ {
	//		if i==0 {
	//			dp[i][k][0] = 0
	//			dp[i][k][1] = -prices[i]
	//			continue
	//		}
	//		dp[i][k][0] = int(math.Max(float64(dp[i-1][k][0]), float64(dp[i-1][k][1]+prices[i])))
	//		dp[i][k][1] = int(math.Max(float64(dp[i-1][k][1]), float64(dp[i-1][k-1][0]-prices[i])))
	//	}
	//}
	//return dp[len(prices)-1][2][0]

	// use variable, change the sequence so that the value wont be change before it is used
	dp10, dp20, dp11, dp21 := 0, 0, math.MinInt64, math.MinInt64
	for i := 0; i < len(prices); i++ {
		dp20 = int(math.Max(float64(dp20), float64(dp21+prices[i])))
		dp21 = int(math.Max(float64(dp21), float64(dp10-prices[i])))
		dp10 = int(math.Max(float64(dp10), float64(dp11+prices[i])))
		dp11 = int(math.Max(float64(dp11), float64(0-prices[i])))
	}
	return dp20
}

// 188
func maxProfitTradeKTimes(k int, prices []int) int {

	// if k is too big, bigger than half of the price array size, means we can do infinite try
	if k > len(prices)/2 {
		return maxProfitInfTimes(prices)
	}
	// normal
	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = int(math.Max(float64(dp[i-1][j][0]), float64(dp[i-1][j][1]+prices[i])))
			dp[i][j][1] = int(math.Max(float64(dp[i-1][j][1]), float64(dp[i-1][j-1][0]-prices[i])))
		}
	}
	return dp[len(prices)-1][k][0]
}

// 309
func maxProfitWithCoolDown(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	b0, b1 := -prices[0], -prices[0]
	s0, s1, s2 := 0, 0, 0
	for i := 1; i < len(prices); i++ {
		b0 = int(math.Max(float64(b1), float64(s2-prices[i])))
		s0 = int(math.Max(float64(s1), float64(b1+prices[i])))
		b1 = b0
		s2 = s1
		s1 = s0
	}
	return s0
}

func maxProfitWithCoolDownsecond(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			// s0, s2
			dp[i][0] = 0
			// b0
			dp[i][1] = -prices[i]
			continue
		}
		if i == 1 {
			// s1
			dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
			// b1
			dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i])))
			continue
		}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-2][0]-prices[i])))
	}
	return dp[len(prices)-1][0]
}

// 722
func maxProfitWithFee(prices []int, fee int) int {
	dp0, dp1 := 0, math.MinInt64

	for i := 0; i < len(prices); i++ {
		temp := dp0
		dp0 = int(math.Max(float64(dp0), float64(dp1+prices[i])))
		dp1 = int(math.Max(float64(dp1), float64(temp-prices[i]-fee)))
	}
	return dp0
}
