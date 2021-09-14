package binaryTree

import (
	"math"
)

// 508
func findFrequentTreeSum(root *TreeNode) []int {
	resMap := make(map[int]int)
	freqDfs(resMap, root)
	res := make([]int, 0)
	max := 0
	for _, v := range resMap {
		max = int(math.Max(float64(max), float64(v)))
	}
	for k, v := range resMap {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func freqDfs(m map[int]int, root *TreeNode) int {
	re := root.Val
	if root.Left != nil {
		re = re + freqDfs(m, root.Left)
	}
	if root.Right != nil {
		re = re + freqDfs(m, root.Right)
	}
	m[re]++
	return re
}
