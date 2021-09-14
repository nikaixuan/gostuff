package binaryTree

import (
	"fmt"
)

// 114
func flatten(root *TreeNode) {
	newSlice := make([]int, 0)
	flatDfs(root, &newSlice)
	fmt.Println(newSlice)
	root = nil
	//root = buildL(newSlice, 0)
}

func flatDfs(ol *TreeNode, slice *[]int) {
	if ol == nil {
		return
	}
	*slice = append(*slice, ol.Val)
	flatDfs(ol.Left, slice)
	flatDfs(ol.Right, slice)
}

func buildL(s []int, n int) *TreeNode {
	if n == len(s) {
		return nil
	}
	root := &TreeNode{}
	root.Val = s[n]
	root.Right = buildL(s, n+1)
	return root
}

// correct answer, pre-order
func flattenDfs(root *TreeNode) {
	if root == nil {
		return
	}
	_ = recDec(root)

}

func recDec(root *TreeNode) *TreeNode {
	right := root.Right
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		// root now is the left child of the last left node(with left)
		root = recDec(root.Right)
	}
	// begin to append right node from bottom to top
	if right != nil {
		root.Right = right
		// root now is possible left/right child of last right node
		root = recDec(root.Right)
	}
	return root
}

// post-order
func flattenPostorder(root *TreeNode) {
	if root == nil {
		return
	}
	flattenPostorder(root.Left)
	flattenPostorder(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = right
}
