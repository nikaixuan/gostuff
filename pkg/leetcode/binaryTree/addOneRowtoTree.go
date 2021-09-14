package binaryTree

// 623
// bfs
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	queue := make([]*TreeNode, 0)
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		depth--
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			if depth == 1 {
				newLeft := &TreeNode{val, curr.Left, nil}
				newRight := &TreeNode{val, nil, curr.Right}
				curr.Left = newLeft
				curr.Right = newRight
			} else {
				if curr.Left != nil {
					queue = append(queue, curr.Left)
				}
				if curr.Right != nil {
					queue = append(queue, curr.Right)
				}
			}
		}
		if depth == 1 {
			break
		}
	}
	return root
}
