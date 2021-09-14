package binaryTree

import (
	"math"
)

// 662
func widthOfBinaryTree(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	iq := make([]int, 0)
	queue = append(queue, root)
	iq = append(iq, 1)
	max := 0
	for len(queue) > 0 {
		l := len(queue)
		left, right := 0, 0
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			index := iq[0]
			iq = iq[1:]
			if i == 0 {
				left = index
			}
			if i == l-1 {
				right = index
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
				iq = append(iq, index*2)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
				iq = append(iq, index*2+1)
			}
		}
		max = int(math.Max(float64(max), float64(right-left+1)))
	}
	return max
}
