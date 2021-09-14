package array

import "math"

// 128
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	res := 0
	for _, v := range nums {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for k, _ := range set {
		if _, ok := set[k-1]; !ok {
			x := k + 1
			for _, ok := set[x]; ok; {
				x = x + 1
				_, ok = set[x]
			}
			res = int(math.Max(float64(res), float64(x-k)))
		}
	}
	return res
}
