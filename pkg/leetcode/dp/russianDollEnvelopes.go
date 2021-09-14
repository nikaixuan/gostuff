package dp

import (
	"sort"
)

// 354
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	size := 0
	height := make([]int, len(envelopes))
	piles := make([]int, len(envelopes))
	for i, e := range envelopes {
		height[i] = e[1]
	}
	for _, h := range height {
		left, right := 0, size
		for left < right {
			mid := (left + right) / 2
			if piles[mid] < h {
				left = mid + 1
			} else {
				right = mid
			}
		}
		piles[left] = h
		if left == size {
			size++
		}
	}
	return size
}
