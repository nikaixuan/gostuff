package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 701
func insertIntoBST(root *binaryTree.TreeNode, val int) *binaryTree.TreeNode {
	if root == nil {
		return &binaryTree.TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
