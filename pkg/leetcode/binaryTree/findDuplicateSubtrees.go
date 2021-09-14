package binaryTree

import (
	"strconv"
)

// 652
var (
	res []*TreeNode
	set map[string]int
)

// serialize the binary tree
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res = make([]*TreeNode, 0)
	set = make(map[string]int)
	_ = findDuplicate(root)
	return res
}

func findDuplicate(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := findDuplicate(root.Left)
	right := findDuplicate(root.Right)
	curr := left + "," + right + "," + strconv.Itoa(root.Val)
	if v, ok := set[curr]; ok {
		if v == 1 {
			res = append(res, root)
		}
		set[curr]++
	} else {
		set[curr] = 1
	}
	return curr
}
