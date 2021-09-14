package backtracking

import "sort"

// must sort first
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	r := make([]int, 0)
	sort.Ints(nums)
	subsetWithDupBack(&res, &r, nums, 0)
	return res
}

func subsetWithDupBack(res *[][]int, r *[]int, nums []int, n int) {
	temp := make([]int, len(*r))
	copy(temp, *r)
	*res = append(*res, temp)
	for i := n; i < len(nums); i++ {
		if i > n && nums[i] == nums[i-1] {
			continue
		}
		*r = append(*r, nums[i])
		subsetWithDupBack(res, r, nums, i+1)
		*r = (*r)[:len(*r)-1]
	}
}
