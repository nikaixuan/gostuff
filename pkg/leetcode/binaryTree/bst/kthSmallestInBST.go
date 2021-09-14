package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 230
func kthSmallest(root *binaryTree.TreeNode, k int) int {
	res := make([]int, 0)
	inOrder(root, k, &res)
	return res[k-1]
}

func inOrder(root *binaryTree.TreeNode, k int, n *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, k, n)
	*n = append(*n, root.Val)
	inOrder(root.Right, k, n)
}
