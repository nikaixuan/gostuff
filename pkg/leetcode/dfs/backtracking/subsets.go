package backtracking

// 78
// no need to sort
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	r := make([]int, 0)
	subsetBack(&res, r, nums, 0)
	return res
}

func subsetBack(res *[][]int, r []int, nums []int, n int) {
	// 用temp拷贝数组，以防止已加入结果的数组被修改
	temp := make([]int, 0)
	*res = append(*res, append(temp, r...))
	for i := n; i < len(nums); i++ {
		r = append(r, nums[i])
		subsetBack(res, r, nums, i+1)
		r = r[:len(r)-1]
	}
}
