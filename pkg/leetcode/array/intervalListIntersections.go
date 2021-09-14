package array

import (
	"math"
	"sort"
)

// 986
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// 因为同一个list是disjoint的，所以同一list的数组不会有交集，可以接起来找交集
	allList := append(firstList, secondList...)
	// 排序
	sort.Slice(allList, func(i, j int) bool {
		if allList[i][0] == allList[j][0] {
			return allList[i][1] > allList[j][1]
		}
		return allList[i][0] < allList[j][0]
	})
	res := make([][]int, 0)
	right := allList[0][1]
	for i := 1; i < len(allList); i++ {
		inv := allList[i]
		// 会有三种情况，1-有交集并被包含，2-不包含但有交集，3-无交集
		// right可作为记录前面最大右界的变量
		// 如果新的数组左界小于right，说明有交集
		if inv[0] <= right {
			// 交集左界是新数组的左界，右界是原最大右界和新右界的最小值（因为取交集），包含了1，2两种情况
			res = append(res, []int{inv[0], int(math.Min(float64(inv[1]), float64(right)))})
		}
		// 第3种情况无交集，但需要更新最大右界用来找后面的交集，而1，2种情况同样需要更新最大右界
		// 左界不需要记录，因为左界唯一的作用是和最大右界比较找交集，如果没有交集，推进最大右边界就好
		right = int(math.Max(float64(inv[1]), float64(right)))
	}
	return res
}

// mergesort
func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		start := int(math.Max(float64(firstList[i][0]), float64(secondList[j][0])))
		end := int(math.Min(float64(firstList[i][1]), float64(secondList[j][1])))
		if start <= end {
			res = append(res, []int{start, end})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
