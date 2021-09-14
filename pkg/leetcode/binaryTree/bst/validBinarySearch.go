package bst

import (
	"fmt"
	"goproject/pkg/leetcode/binaryTree"
)

// 98
func isValidBST(root *binaryTree.TreeNode) bool {
	arr := make([]int, 0)
	bstDfs(&arr, root)
	if len(arr) == 1 {
		return true
	}
	// O(n), not good, better to have O(logn)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func bstDfs(a *[]int, root *binaryTree.TreeNode) {
	if root == nil {
		return
	}
	bstDfs(a, root.Left)
	fmt.Println(root.Val)
	*a = append(*a, root.Val)
	bstDfs(a, root.Right)
}

// Better answer
// RecValidate(node, nil, nil)
func recValidate(n, min, max *binaryTree.TreeNode) bool {
	if n == nil {
		return true
	}
	if min != nil && n.Val <= min.Val {
		return false
	}
	if max != nil && n.Val >= max.Val {
		return false
	}
	return recValidate(n.Left, min, n) && recValidate(n.Right, n, max)
}
