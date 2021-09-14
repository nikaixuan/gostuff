package designDataStructure

// 496
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make(map[int]int)
	r := make([]int, 0, len(nums1))
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[nums2[i]] = stack[len(stack)-1]
		} else {
			res[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	for i := range nums1 {
		r = append(r, res[nums1[i]])
	}
	return r
}
