package bst

import (
	"goproject/pkg/leetcode/binaryTree"
	"math"
)

// 1373
var max int

func maxSumBST(root *binaryTree.TreeNode) int {
	_, _, _, _ = traverMaxsum(root)
	return int(math.Max(0, float64(max)))
}

func traverMaxsum(root *binaryTree.TreeNode) (int, int, int, bool) {
	if root.Right == nil && root.Left == nil {
		max = int(math.Max(float64(root.Val), float64(max)))
		return root.Val, root.Val, root.Val, true
	}
	if root.Right == nil {
		left, maxl, minl, isleft := traverMaxsum(root.Left)
		if isleft && root.Val > maxl {
			max = int(math.Max(float64(left+root.Val), float64(max)))
			return left + root.Val, root.Val, minl, true
		}
		return 0, 0, 0, false
	}
	if root.Left == nil {
		right, maxr, minr, isright := traverMaxsum(root.Right)
		if isright && root.Val < minr {
			max = int(math.Max(float64(right+root.Val), float64(max)))
			return right + root.Val, maxr, root.Val, true
		}
		return 0, 0, 0, false
	}
	left, maxl, minl, isleft := traverMaxsum(root.Left)
	right, maxr, minr, isright := traverMaxsum(root.Right)
	if isleft && isright && root.Val > maxl && root.Val < minr {
		max = int(math.Max(float64(max), float64(left+right+root.Val)))
		return left + right + root.Val, maxr, minl, true
	}
	return 0, 0, 0, false
}
