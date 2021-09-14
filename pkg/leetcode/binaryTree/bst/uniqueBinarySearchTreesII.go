package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 95
func generateTrees(n int) []*binaryTree.TreeNode {
	return buildBSTFromNumber(1, n)
}

func buildBSTFromNumber(start, end int) []*binaryTree.TreeNode {
	res := make([]*binaryTree.TreeNode, 0)
	if start > end {
		// return an empty slice so the parent append func will not have nil pointer
		return append(res, nil)
	}
	// from start to end, every value needs to be root once
	for i := start; i <= end; i++ {
		// need to exclude i because it needs to be the root value
		left := buildBSTFromNumber(start, i-1)
		right := buildBSTFromNumber(i+1, end)
		// permutation
		for _, l := range left {
			for _, r := range right {
				root := &binaryTree.TreeNode{Val: i}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}
