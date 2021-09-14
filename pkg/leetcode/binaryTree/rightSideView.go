package binaryTree

// 199
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	viewDfs(root, &result, 0)
	return result
}

func viewDfs(r *TreeNode, res *[]int, curr int) {
	if r == nil {
		return
	}
	if len(*res) == curr {
		*res = append(*res, r.Val)
	}
	viewDfs(r.Right, res, curr+1)
	viewDfs(r.Left, res, curr+1)
}
