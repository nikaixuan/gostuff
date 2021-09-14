package dp

import "math"

// 514
func findRotateSteps(ring string, key string) int {
	m := make(map[byte][]int)
	for i := range ring {
		if _, ok := m[ring[i]]; !ok {
			m[ring[i]] = []int{i}
		} else {
			curr := m[ring[i]]
			m[ring[i]] = append(curr, i)
		}
	}
	dp := make([][]int, len(key)+1)
	for i := range dp {
		dp[i] = make([]int, len(ring))
	}
	for i := len(key) - 1; i >= 0; i-- {
		for j := 0; j < len(ring); j++ {
			dp[i][j] = math.MaxInt64
			for _, v := range m[key[i]] {
				diff := int(math.Abs(float64(j - v)))
				dis := int(math.Min(float64(diff), float64(len(ring)-diff)))
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i+1][v]+dis)))
			}
		}
	}
	return dp[0][0] + len(key)
}
