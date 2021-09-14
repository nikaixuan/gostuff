package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 538 1038
var sum int

func convertBST(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	sum = 0
	traverseBSTGetSlice(root)
	return root
}

func traverseBSTGetSlice(root *binaryTree.TreeNode) {
	if root == nil {
		return
	}
	traverseBSTGetSlice(root.Right)
	sum += root.Val
	root.Val = sum
	traverseBSTGetSlice(root.Left)
}
