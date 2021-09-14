package binaryTree

// 865
// 1123
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res, _ := postOrderAllDeepest(root)
	return res
}

func postOrderAllDeepest(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, -1
	}
	leftnode, leftdeep := postOrderAllDeepest(root.Left)
	rightnode, rightdeep := postOrderAllDeepest(root.Right)
	if leftdeep < rightdeep {
		// return the node with most subtree from child and record the deep value
		return rightnode, rightdeep + 1
	} else if leftdeep > rightdeep {
		return leftnode, leftdeep + 1
	} else {
		// when left deep == right deep, current node is the node that have most subtree
		return root, leftdeep + 1
	}
}
