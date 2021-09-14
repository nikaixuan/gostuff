package binaryTree

// 105
func buildTree(preorder []int, inorder []int) *TreeNode {
	return constructTreeHelper(0, 0, len(inorder)-1, preorder, inorder)
}

func constructTreeHelper(preStart, inStart, inEnd int, p, i []int) *TreeNode {
	if preStart > len(p)-1 || inStart > inEnd {
		return nil
	}
	treeNode := &TreeNode{Val: p[preStart]}
	iIndex := 0
	for j := inStart; j <= inEnd; j++ {
		if p[preStart] == i[j] {
			iIndex = j
		}
	}
	treeNode.Left = constructTreeHelper(preStart+1, inStart, iIndex-1, p, i)
	treeNode.Right = constructTreeHelper(preStart+iIndex-inStart+1, iIndex+1, inEnd, p, i)
	return treeNode
}
