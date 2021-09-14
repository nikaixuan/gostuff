package binaryTree

// 894
// similar to build all bst
func allPossibleFBT(n int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if n == 1 {
		return append(res, &TreeNode{Val: 0})
	} else {
		for i := 1; i < n; i += 2 {
			left := allPossibleFBT(i)
			right := allPossibleFBT(n - i - 1)
			for _, l := range left {
				for _, r := range right {
					root := &TreeNode{Val: 0}
					root.Left = l
					root.Right = r
					res = append(res, root)
				}
			}
		}
	}
	return res
}
