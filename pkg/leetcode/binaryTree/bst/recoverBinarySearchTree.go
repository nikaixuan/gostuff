package bst

import (
	"goproject/pkg/leetcode/binaryTree"
	"math"
)

func recoverTree(root *binaryTree.TreeNode) {
	if root == nil {
		return
	}
	var f, s *binaryTree.TreeNode
	f, s = nil, nil
	prev := &binaryTree.TreeNode{Val: math.MinInt64}
	recoverDfs(root, &prev, &f, &s)
	if f != nil && s != nil {
		f.Val, s.Val = s.Val, f.Val
	}
}

func recoverDfs(r *binaryTree.TreeNode, prev, first, second **binaryTree.TreeNode) {
	if r == nil {
		return
	}
	recoverDfs(r.Left, prev, first, second)
	if *first == nil && (*prev).Val >= r.Val {
		*first = *prev
	}
	if *first != nil && (*prev).Val >= r.Val {
		*second = r
	}
	*prev = r
	recoverDfs(r.Right, prev, first, second)
}
