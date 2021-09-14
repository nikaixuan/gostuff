package binaryTree

import (
	"strconv"
)

// 129
func sumNumbers(root *TreeNode) int {
	res := make([]string, 0)
	result := 0
	sumNDfs(&res, "", root)
	for i := range res {
		ii, err := strconv.Atoi(res[i])
		if err == nil {
			result += ii
		}
	}
	return result
}

func sumNDfs(res *[]string, r string, root *TreeNode) {
	newR := r + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, r)
		return
	}
	if root.Left != nil {
		sumNDfs(res, newR, root.Left)
	}
	if root.Right != nil {
		sumNDfs(res, newR, root.Right)
	}
}
