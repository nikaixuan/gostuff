package binarySearch

import "math"

func search1(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for high > low {
		mid := low + (high-low)/2
		num := 0
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			num = nums[mid]
		} else {
			if target < nums[0] {
				num = math.MinInt64
			} else {
				num = math.MaxInt64
			}
		}
		if num < target {
			low = mid + 1
		} else if num > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
