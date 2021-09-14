package backtracking

import "sort"

// 47
func permutation2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	var r []int
	sort.Ints(nums)
	visited := make(map[int]bool)
	perBackTrack2(nums, &res, r, visited)
	return res
}

func perBackTrack2(l []int, res *[][]int, r []int, v map[int]bool) {
	if len(r) == len(l) {
		var rr []int
		*res = append(*res, append(rr, r...))

		//temp := make([]int, len(r))
		//copy(temp, r) //copy list into temp, so it won't impact subsequent process on list
		//*res = append(*res, temp)
		return
	}
	for i := 0; i < len(l); i++ {
		if v[i] {
			continue
		}
		// 本质：不打乱连续数字的顺序，让连续数字组合只出现一次
		if i > 0 && l[i] == l[i-1] && !v[i-1] {
			continue
		}
		r = append(r, l[i])
		v[i] = true
		perBackTrack2(l, res, r, v)
		r = r[:len(r)-1]
		v[i] = false
	}
}
