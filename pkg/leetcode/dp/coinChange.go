package dp

import (
	"math"
)

// 322
func coinChange(coins []int, amount int) int {
	queue := make([]int, 0)
	m := make(map[int]struct{})
	res := 0
	queue = append(queue, amount)
	m[amount] = struct{}{}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == 0 {
				return res
			}
			for c := range coins {
				next := curr - coins[c]
				if _, ok := m[next]; !ok && next >= 0 {
					queue = append(queue, next)
					m[next] = struct{}{}
				}
			}
		}
		res = res + 1
	}
	return -1
}

// dp
func coinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		// equal to use math.max
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 与coinchange2相比，遍历方法不同，dp定义不同
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			// 通过dp[i-coin]凑出的面值+1 凑出当前面值
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
