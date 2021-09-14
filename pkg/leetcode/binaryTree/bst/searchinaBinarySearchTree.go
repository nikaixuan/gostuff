package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 700
func searchBST(root *binaryTree.TreeNode, val int) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return nil
}
