package array

import (
	"sort"
)

// 15
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res1 := map[*[]int]struct{}{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := i + 1
		j := n - 1
		for k < j {
			if nums[i]+nums[k]+nums[j] == 0 {
				res1[&[]int{nums[i], nums[k], nums[j]}] = struct{}{}
				k++
				j--
				for k < j && nums[k] == nums[k-1] {
					k++
				}
				for k < j && nums[j] == nums[j+1] {
					j--
				}
			} else if nums[i]+nums[k]+nums[j] < 0 {
				k++
			} else if nums[i]+nums[k]+nums[j] > 0 {
				j--
			}
		}
	}
	var res [][]int
	for k, _ := range res1 {
		res = append(res, *k)
	}
	return res

}
