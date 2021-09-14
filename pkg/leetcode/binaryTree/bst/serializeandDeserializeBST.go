package bst

import (
	"goproject/pkg/leetcode/binaryTree"
	"math"
)

type Codec struct {
}

func codecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *binaryTree.TreeNode) string {
	s := make([]rune, 0)
	serDfs(root, &s)
	return string(s)
}

func serDfs(root *binaryTree.TreeNode, s *[]rune) {
	if root == nil {
		return
	}
	*s = append(*s, rune(root.Val))
	serDfs(root.Left, s)
	serDfs(root.Right, s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *binaryTree.TreeNode {
	if data == "" {
		return nil
	}
	queue := []rune(data)
	return derDfs(&queue, math.MinInt64, math.MaxInt64)
}

func derDfs(queue *[]rune, low int, upper int) *binaryTree.TreeNode {
	if len(*queue) == 0 {
		return nil
	}
	curr := (*queue)[0]
	v := int(curr)
	if v < low || v > upper {
		return nil
	}
	*queue = (*queue)[1:]
	treenode := &binaryTree.TreeNode{Val: v}
	treenode.Left = derDfs(queue, low, v)
	treenode.Right = derDfs(queue, v, upper)
	return treenode
}
