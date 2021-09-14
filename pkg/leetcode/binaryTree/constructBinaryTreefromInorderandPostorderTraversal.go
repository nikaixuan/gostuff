package binaryTree

// 106
func buildTreeInorderPostorder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	currRootVal := postorder[len(postorder)-1]
	node := &TreeNode{Val: currRootVal}
	index := 0
	for i := range inorder {
		if inorder[i] == currRootVal {
			index = i
		}
	}
	node.Left = buildTreeInorderPostorder(inorder[:index], postorder[:index])
	if index == len(inorder)-1 {
		node.Right = nil
	} else {
		node.Right = buildTreeInorderPostorder(inorder[index+1:], postorder[index:len(postorder)-1])
	}
	return node
}
