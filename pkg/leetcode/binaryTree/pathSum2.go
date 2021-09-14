package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 113
func pathSum2(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	res := make([]int, 0)
	pathSumDfs(root, targetSum, res, &result)
	return result
}

func pathSumDfs(node *TreeNode, t int, res []int, result *[][]int) {
	if node == nil {
		return
	}
	if t-node.Val == 0 && node.Left == nil && node.Right == nil {
		temp := make([]int, 0)
		temp = append(temp, res...)
		*result = append(*result, append(temp, node.Val))
		return
	}
	pathSumDfs(node.Left, t-node.Val, append(res, node.Val), result)
	pathSumDfs(node.Right, t-node.Val, append(res, node.Val), result)
}
