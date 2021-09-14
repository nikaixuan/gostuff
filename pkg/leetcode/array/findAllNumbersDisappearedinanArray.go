package array

import (
	"math"
)

// 448
// 联系索引和数字，很巧妙
func findDisappearedNumbers(nums []int) []int {
	// make nums[i] as index and change that index number to < 0
	// when iteration ends, index of number still >0 is the missing number
	for i := range nums {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	res := make([]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

// 645
func findErrorNums(nums []int) []int {
	res := make([]int, 0)
	for i := range nums {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] < 0 {
			res = append(res, index+1)
		}
		nums[index] = -nums[index]
	}
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
