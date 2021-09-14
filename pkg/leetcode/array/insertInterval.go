package array

import "math"

// 57
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < len(intervals) && newInterval[0] < intervals[i][1] {
		newInterval[0] = int(math.Min(float64(newInterval[0]), float64(intervals[i][0])))
		newInterval[1] = int(math.Max(float64(newInterval[1]), float64(intervals[i][1])))
	}
	res = append(res, newInterval)
	i++
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}
