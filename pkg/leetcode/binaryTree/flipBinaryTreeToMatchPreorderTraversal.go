package binaryTree

var index int

// 971
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	index = 0
	res := make([]int, 0)
	if poFlip(root, voyage, &res) {
		return res
	}
	return []int{-1}
}

func poFlip(root *TreeNode, voyage []int, res *[]int) bool {

	if root == nil {
		return true
	}
	// one of the base case
	// check the value
	if root.Val != voyage[index] {
		return false
	}
	index++
	// check whether to flip
	if root.Left != nil && root.Left.Val != voyage[index] {
		*res = append(*res, root.Val)
		return poFlip(root.Right, voyage, res) && poFlip(root.Left, voyage, res)
	}
	// no flip
	return poFlip(root.Left, voyage, res) && poFlip(root.Right, voyage, res)
}
