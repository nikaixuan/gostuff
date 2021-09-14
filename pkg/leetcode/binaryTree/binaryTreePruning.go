package binaryTree

// 814
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := pruneTree(root.Left)
	right := pruneTree(root.Right)
	if left == nil && right == nil && root.Val == 0 {
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}
