package greedy

import "math"

// 45
// greedy is better
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i+j] = int(math.Min(float64(dp[i+j]), float64(dp[i]+1)))
		}
	}
	return dp[len(nums)-1]
}

// greedy
func jumpGreedy(nums []int) int {
	count := 0
	if len(nums) == 1 {
		return count
	}
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			l := i - j
			if nums[j] >= l && l > 1 {
				i = i - l + 1
				break
			}
		}
		count++
	}
	return count
}

func jumpGreedy2(nums []int) int {
	count := 0
	if len(nums) == 1 {
		return count
	}
	maxFar := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxFar = int(math.Max(float64(maxFar), float64(i+nums[i])))
		if end == i {
			end = maxFar
			count++
		}
	}
	return count
}
