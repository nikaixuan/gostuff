package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 450
func deleteNode(root *binaryTree.TreeNode, key int) *binaryTree.TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Left
		}
		if root.Right == nil {
			return root.Right
		}
		// use min val in right tree or max value in right tree to replace itself
		// find the node
		min := getMin(root.Right)
		// switch
		root.Val = min
		// delete that node
		root.Right = deleteNode(root.Right, min)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(root *binaryTree.TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return getMin(root.Left)
}
