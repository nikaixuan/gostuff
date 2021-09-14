package binaryTree

import (
	"strconv"
	"strings"
)

// 297
type btCodec struct {
}

func btCodecConstructor() btCodec {
	return btCodec{}
}

// Serializes a tree to a single string. bfs solution
func (this *btCodec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := ""
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				res = res + "n "
				continue
			}
			res = res + strconv.Itoa(curr.Val) + " "
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		}
	}
	return strings.TrimRight(res, " ")
}

// Deserializes your encoded data to tree.
func (this *btCodec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	queue := make([]*TreeNode, 0)
	nodes := strings.Split(data, " ")
	v, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: v}
	queue = append(queue, root)
	for i := 1; i < len(nodes); i++ {
		parent := queue[0]
		queue = queue[1:]
		if nodes[i] != "n" {
			l, _ := strconv.Atoi(nodes[i])
			left := &TreeNode{Val: l}
			parent.Left = left
			queue = append(queue, left)
		}
		i++
		if nodes[i] != "n" {
			r, _ := strconv.Atoi(nodes[i])
			right := &TreeNode{Val: r}
			parent.Right = right
			queue = append(queue, right)
		}
	}
	return root
}

// pre-order dfs solution
func (this *btCodec) serializepre(root *TreeNode) string {

	if root == nil {
		return "#"
	}
	res := ""
	serializetrav(root, &res)
	return res
}

func serializetrav(root *TreeNode, res *string) {
	if root == nil {
		*res = *res + "#" + ","
		return
	}
	*res = *res + strconv.Itoa(root.Val) + ","
	serializetrav(root.Left, res)
	serializetrav(root.Right, res)
}

func (this *btCodec) deserializepre(data string) *TreeNode {
	s := strings.Split(data, ",")
	return destrav(&s)
}

func destrav(s *[]string) *TreeNode {
	ss := (*s)[0]
	if len(*s) == 1 {
		return nil
	}
	*s = (*s)[1:]
	if ss == "#" {
		return nil
	}
	i, _ := strconv.Atoi(ss)
	node := &TreeNode{Val: i}
	node.Left = destrav(s)
	node.Right = destrav(s)
	return node
}
