package binarySearch

import (
	"sort"
)

func smallestDivisor(nums []int, threshold int) int {
	sort.Ints(nums)
	left, right := 1, nums[len(nums)-1]
	for left < right {
		mid := left + (right-left)/2
		// need to think about is there multiple mid can lead to a same r, if so, do not return the first mid, squeeze to the left or right
		r := divide(nums, mid)
		if r == threshold {
			right = mid
		} else if r > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// the smaller the number, the bigger the threshold
func divide(nums []int, num int) int {
	res := 0
	for _, n := range nums {
		res += n / num
		if n%num > 0 {
			res++
		}
	}
	return res
}
