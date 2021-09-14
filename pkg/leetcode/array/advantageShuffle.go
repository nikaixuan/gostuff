package array

import "sort"

// 870
func advantageCount(nums1 []int, nums2 []int) []int {
	n2 := make([][2]int, 0)
	for i, v := range nums2 {
		n2 = append(n2, [2]int{i, v})
	}
	sort.Ints(nums1)
	sort.Slice(n2, func(i, j int) bool {
		return n2[i][1] < n2[j][1]
	})
	res := make([]int, len(n2))
	i, j := 0, len(nums1)-1
	for k := len(n2) - 1; k >= 0; k-- {
		if n2[k][1] >= nums1[j] {
			res[n2[k][0]] = nums1[i]
			i++
		} else {
			res[n2[k][0]] = nums1[j]
			j--
		}
	}
	return res
}
