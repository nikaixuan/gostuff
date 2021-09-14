package dp

import (
	"math"
	"sort"
)

// 1235
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	arr := make([][]int, len(startTime))
	for i := range arr {
		arr[i] = make([]int, 3)
		arr[i][0] = startTime[i]
		arr[i][1] = endTime[i]
		arr[i][2] = profit[i]
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	dp := make([]int, len(startTime))
	dp[len(dp)-1] = arr[len(dp)-1][2]
	for i := len(dp) - 2; i >= 0; i-- {
		// find whether there is job start with current end time
		next_index := nextIndex(i, arr)
		if next_index != -1 {
			// found [1,3] + [3,6] compare with prev profit
			dp[i] = int(math.Max(float64(dp[i+1]), float64(arr[i][2]+dp[next_index])))
		} else {
			// not found, current profit compare with prev profit
			dp[i] = int(math.Max(float64(dp[i+1]), float64(arr[i][2])))
		}
	}
	return dp[0]
}

func nextIndex(k int, jobs [][]int) int {
	left, right := k+1, len(jobs)-1
	for left <= right {
		mid := left + (right-left)/2
		if jobs[mid][0] >= jobs[k][1] {
			if jobs[mid-1][0] >= jobs[k][1] {
				right = mid - 1
			} else {
				return mid
			}
		} else {
			left = mid + 1
		}
	}
	return -1

	// linear search
	//
	//for i:=k+1;i<len(jobs);i++ {
	//	if jobs[k][1]<=jobs[i][0] {
	//		return i
	//	}
	//}
	//return -1
}
