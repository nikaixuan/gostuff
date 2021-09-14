package dp

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	// store the index of the divide chain
	index := make([]int, len(nums))

	max := 0

	// max divide index
	ii := -1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		index[i] = -1
		// traverse from back to current and find the divide chain
		for j := i - 1; j >= 0; j-- {
			// dp[j]+1>dp[i] for eliminate duplicate, (1,2,3,6) will always be 3
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				index[i] = j
			}
		}
		// find max divide and get the index of the last number in the chain
		if dp[i] > max {
			max = dp[i]
			ii = i
		}
	}
	res := make([]int, 0)
	// find the chain using index array
	for ii != -1 {
		res = append(res, nums[ii])
		ii = index[ii]
	}
	return res

}
