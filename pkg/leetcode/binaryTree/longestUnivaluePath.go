package binaryTree

import (
	"math"
)

// 687
var max int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	pOrderCount(root, root.Val)
	return max
}

func pOrderCount(root *TreeNode, value int) int {
	if root == nil {
		return 0
	}
	cl := pOrderCount(root.Left, root.Val)
	cr := pOrderCount(root.Right, root.Val)
	nc := cl + cr
	// if l/r have same value as current node, will return not 0, but the bigger path of it
	// no need to check, directly compare it with max
	max = int(math.Max(float64(max), float64(nc)))
	if value == root.Val {
		return int(math.Max(float64(cl), float64(cr))) + 1
	}
	return 0
}
