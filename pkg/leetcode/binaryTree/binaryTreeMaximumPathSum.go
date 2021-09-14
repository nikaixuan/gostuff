package binaryTree

import (
	"math"
)

var answer = math.MinInt64

func maxPathSum(root *TreeNode) int {
	answer = math.MinInt64
	traversePath(root)
	return answer
}

func traversePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := int(math.Max(0, float64(traversePath(root.Left))))
	right := int(math.Max(0, float64(traversePath(root.Right))))
	answer = int(math.Max(float64(answer), float64(left+right+root.Val)))
	return int(math.Max(float64(left), float64(right))) + root.Val
}
