package designDataStructure

// 503
func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(nums))
	l := len(nums)
	for i := l*2 - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%l] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i%l] = stack[len(stack)-1]
		} else {
			res[i%l] = -1
		}
		stack = append(stack, nums[i%l])
	}
	return res
}
