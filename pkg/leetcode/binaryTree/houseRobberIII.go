package binaryTree

import (
	"math"
)

// 337
func rob(root *TreeNode) int {
	res := robDfs(root)
	return int(math.Max(float64(res[0]), float64(res[1])))
}

func robDfs(root *TreeNode) []int {
	if root == nil {
		return make([]int, 2)
	}
	left := robDfs(root.Left)
	right := robDfs(root.Right)
	res := make([]int, 2)
	res[0] = int(math.Max(float64(left[0]), float64(left[1])) + math.Max(float64(right[0]), float64(right[1])))
	res[1] = left[0] + right[0] + root.Val
	return res
}
