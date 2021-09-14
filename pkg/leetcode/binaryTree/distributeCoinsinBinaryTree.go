package binaryTree

import "math"

// 979
var coinres int

func distributeCoins(root *TreeNode) int {
	coinres = 0
	postOrderCoin(root)
	return coinres
}

func postOrderCoin(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := postOrderCoin(root.Left), postOrderCoin(root.Right)
	move := left + right + root.Val - 1
	coinres += int(math.Abs(float64(move)))
	return move
}
