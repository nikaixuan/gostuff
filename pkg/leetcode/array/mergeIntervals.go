package array

import "sort"

// 56
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	left, right := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		inv := intervals[i]
		if inv[0] <= right && inv[1] > right {
			right = inv[1]
		} else if inv[0] > right {
			res = append(res, []int{left, right})
			left = inv[0]
			right = inv[1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
