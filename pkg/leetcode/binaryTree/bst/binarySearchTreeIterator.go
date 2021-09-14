package bst

import (
	"goproject/pkg/leetcode/binaryTree"
)

// 173
type BSTIterator struct {
	nodeList []int
	index    int
}

func Constructor(root *binaryTree.TreeNode) BSTIterator {
	nl := make([]int, 0)
	treeInorder(root, &nl)
	return BSTIterator{
		nodeList: nl,
		index:    0,
	}
}

func treeInorder(root *binaryTree.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	treeInorder(root.Left, res)
	*res = append(*res, root.Val)
	treeInorder(root.Right, res)
}

func (this *BSTIterator) Next() int {
	curr := this.nodeList[this.index]
	this.index++
	return curr
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.nodeList)
}
