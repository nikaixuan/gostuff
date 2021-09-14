package backtracking

// 39
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var r []int
	if len(candidates) == 0 {
		return result
	}
	// 可不排序
	//sort.Ints(candidates)
	sumBackTracking(candidates, target, &result, r, 0)
	return result
}

func sumBackTracking(candidates []int, target int, result *[][]int, r []int, n int) {
	// base case
	if target == 0 {
		var rr []int
		*result = append(*result, append(rr, r...))
		return
	}
	for i := n; i < len(candidates); i++ {
		// 负数检查可放在循环内也可作为回溯的base case之一，但放在这时间复杂度好一些
		if target-candidates[i] < 0 {
			continue
		}
		sumBackTracking(candidates, target-candidates[i], result, append(r, candidates[i]), i)
	}
}
