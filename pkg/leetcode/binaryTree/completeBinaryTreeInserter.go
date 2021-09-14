package binaryTree

import "math"

type CBTInserter struct {
	size int
	root *TreeNode
}

// 919
func ConstructorCBTInserter(root *TreeNode) CBTInserter {
	return CBTInserter{
		size: countCTBNode(root),
		root: root,
	}
}

func countCTBNode(root *TreeNode) int {
	hl, hr := 0, 0
	l, r := root, root
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	if hr == hl {
		return int(math.Pow(2, float64(hl)) - 1)
	}
	return 1 + countCTBNode(root.Left) + countCTBNode(root.Right)
}

func (this *CBTInserter) Insert(val int) int {
	newSize := this.size + 1
	dir := make([]byte, 0)
	for newSize != 1 {
		if newSize%2 == 0 {
			dir = append(dir, 'l')
		} else {
			dir = append(dir, 'r')
		}
		newSize = newSize / 2
	}
	addRoot := this.root
	for i := len(dir) - 1; i >= 1; i-- {
		if dir[i] == 'l' {
			addRoot = addRoot.Left
		} else {
			addRoot = addRoot.Right
		}
	}
	if dir[0] == 'l' {
		addRoot.Left = &TreeNode{Val: val}
	} else {
		addRoot.Right = &TreeNode{Val: val}
	}
	this.size++
	return addRoot.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
