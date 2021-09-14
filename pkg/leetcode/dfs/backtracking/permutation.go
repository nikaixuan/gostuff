package backtracking

// 46
func permute(list []int) [][]int {
	var res [][]int
	var r []int
	perBackTrack(list, &res, r)
	return res
}

func perBackTrack(l []int, res *[][]int, r []int) {
	if len(r) == len(l) {
		*res = append(*res, r)
	} else {
		for i := 0; i < len(l); i++ {
			if isIn(r, l[i]) {
				continue
			}
			// 不修改现有数组的话就不用撤销选择
			newr := append(r, l[i])
			perBackTrack(l, res, newr)
		}
	}
}

//o(n)复杂度，也可以用set
func isIn(arr []int, n int) bool {
	for i := range arr {
		if arr[i] == n {
			return true
		}
	}
	return false
}
