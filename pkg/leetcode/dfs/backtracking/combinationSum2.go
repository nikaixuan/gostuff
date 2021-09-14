package backtracking

import "sort"

// 40
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var res []int
	sort.Ints(candidates)
	subTwoBacktrack(candidates, target, res, &result, 0)
	return result
}

func subTwoBacktrack(candidates []int, target int, r []int, res *[][]int, n int) {
	if target== 0 {
		temp := append([]int{}, r...)
		*res = append(*res, temp)
		return
	}

	for i:=n;i<len(candidates);i++ {
		if target - candidates[i]<0 {
			continue
		}
		if i>n && candidates[i] == candidates[i-1] {
			continue
		}
		r = append(r, candidates[i])
		subTwoBacktrack(candidates, target - candidates[i], r, res, i+1)
		r = r[:len(r)-1]
	}
}


func combinationSum22(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	var res []int
	subTwoBacktrack22(candidates, target, res, &result)
	return result
}

func subTwoBacktrack22(candidates []int, target int, r []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, r)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			return
		}
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		subTwoBacktrack22(candidates[i+1:], target-candidates[i], append(r, candidates[i]), res)
	}
}

