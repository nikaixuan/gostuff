package binaryTree

import (
	"math"
)

// 222
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	n := 1
	left := root
	for {
		left = left.Left
		if left != nil {
			n = n + 1
		} else {
			break
		}
	}
	if n == 1 {
		return 1
	}
	l := countLast(root, n-2)
	return int(math.Pow(2, float64(n-1))) - 1 + l
}

func countLast(root *TreeNode, l int) int {
	if l == 0 {
		if root.Right != nil && root.Left != nil {
			return 2
		}
		if root.Right != nil || root.Left != nil {
			return 1
		}
		return 0
	}
	return countLast(root.Left, l-1) + countLast(root.Right, l-1)
}

func countNodes2(root *TreeNode) int {
	l, r := root, root
	hl, hr := 0, 0
	for r != nil {
		hr++
		r = r.Right
	}
	for l != nil {
		hl++
		l = l.Left
	}
	if hl == hr {
		return int(math.Pow(2, float64(hr)) - 1)
	}
	return 1 + countNodes2(root.Left) + countNodes2(root.Right)
}
