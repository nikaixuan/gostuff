package array

import "sort"

// 1288
func removeCoveredIntervals(intervals [][]int) int {
	// 相同左界时按右界大到小排序， 后面的算法好考虑
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[j][1] < intervals[i][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	// 代表会被其他区间包含的
	res := 0
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 左右都在已有界内，被包含+1，此left相关检查也可以换成<=right，因为只是检查是否在已有界内
		if intervals[i][0] >= left && intervals[i][1] <= right {
			res++
			// 左界在界内，右界在界外，更新右边
		} else if intervals[i][0] <= right && intervals[i][1] > right {
			right = intervals[i][1]
			// 新区间左界在之前的右界外，更新左右
			// 更新左右很重要，新区间左右在之前界外说明之前的左右都已经没有意义，只需要记录新左右界就好
		} else if intervals[i][0] >= right {
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	return len(intervals) - res
}
