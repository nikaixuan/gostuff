package backtracking

// 77
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	combineBacktrack(1, n, k, &res, []int{})
	return res
}

func combineBacktrack(start, n, k int, res *[][]int, curr []int) {
	if len(curr) == k {
		// 需要拷贝curr，如果直接添加curr，回溯后面对curr进行修改时，前面已经加入结果的数组也会变
		temp := make([]int, 0)
		*res = append(*res, append(temp, curr...))
		return
	}
	for i := start; i <= n; i++ {
		// 正规写法：
		curr = append(curr, i)
		combineBacktrack(i+1, n, k, res, curr)
		curr = curr[:len(curr)-1]

		// 因为我们要的就是对curr的修改不影响for循环的下一个迭代，所以可以创建新参数来append curr， 也可以直接在方法调用中
		// 传递append(curr, i)，都可以达到同样的目的
	}
}
