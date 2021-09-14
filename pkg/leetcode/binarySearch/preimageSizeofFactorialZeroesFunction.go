package binarySearch

import "math"

// 793
func preimageSizeFZF(k int) int {
	left, right := 0, math.MaxInt64
	l, r := 0, 0
	for left < right {
		mid := left + (right-left)/2
		z := trailingZeroes(mid)
		if z >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	l = left

	left, right = 0, math.MaxInt64
	for left < right {
		mid := left + (right-left)/2
		z := trailingZeroes(mid)
		if z <= k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// for both looking for left and right bound
	// left variable always will be the first number that trailingZeroes(n) larger than k
	// 二分查找压缩边界的左界：所有小于target的元素。右界：所有大于target的元素
	// 所以当元素不存在时，右界会大于左界，所以r-l+1会等于0
	r = left - 1
	return r - l + 1
}

func trailingZeroes(n int) int {
	res := 0
	for ; n > 0; n /= 5 {
		res += n / 5
	}
	return res
}
