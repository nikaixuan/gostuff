package greedy

import (
	"math"
	"sort"
)

// 1024
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	if clips[0][0] != 0 {
		return -1
	}
	currEnd, nextEnd, count, i := 0, 0, 0, 0
	for i < len(clips) && clips[i][0] <= currEnd {
		for i < len(clips) && clips[i][0] <= currEnd {
			nextEnd = int(math.Max(float64(nextEnd), float64(clips[i][1])))
			i++
		}
		count++
		currEnd = nextEnd
		if currEnd >= time {
			return count
		}
	}
	return -1
}
