package binaryTree

// pre-order
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max, index := 0, 0
	for i, n := range nums {
		if max < n {
			max = n
			index = i
		}
	}
	node := &TreeNode{Val: nums[index]}
	node.Left = constructMaximumBinaryTree(nums[:index])
	if index == len(nums)-1 {
		node.Right = nil
	} else {
		node.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return node
}
