package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 669
func trimBST(root *binaryTree.TreeNode, low int, high int) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, root.Val)
	root.Right = trimBST(root.Right, root.Val, high)
	return root
}
