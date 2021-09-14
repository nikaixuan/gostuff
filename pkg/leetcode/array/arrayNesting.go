package array

import "math"

// 565
func arrayNesting(nums []int) int {
	res := 0
	b := make([]bool, len(nums))
	// dfs, check the chain and mark tracked index as true.
	// skip true index because it is visited and already the nest length contains it.
	for i := 0; i < len(nums); i++ {
		count := 0
		for !b[i] {
			b[i] = true
			count++
			i = nums[i]
		}
		res = int(math.Max(float64(res), float64(count)))
	}
	return res
}
