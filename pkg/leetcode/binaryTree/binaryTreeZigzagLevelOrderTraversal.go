package binaryTree

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	zigzagBfs(root, 0, &res)
	return res
}

func zigzagBfs(root *TreeNode, l int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == l {
		*res = append(*res, []int{})
	}
	if l%2 == 1 {
		(*res)[l] = append([]int{root.Val}, (*res)[l]...)
	} else {
		(*res)[l] = append((*res)[l], root.Val)
	}
	zigzagBfs(root.Left, l+1, res)
	zigzagBfs(root.Right, l+1, res)
}
