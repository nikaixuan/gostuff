package dp

import "math"

// 300
// not clear
// binary search solution
func lengthOfLIS(nums []int) int {
	tails := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		i, j := 0, size
		for i != j {
			m := (i + j) / 2
			if tails[m] < num {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = num
		if i == size {
			size++
		}
	}
	return size
}

// dp solution
func lengthOfLISDp(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(1+dp[j])))
			}
		}
	}
	for i := range dp {
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}
