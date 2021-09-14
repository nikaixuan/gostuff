package binaryTree

// 513
// bfs
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	res := root.Val
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if i == 0 {
				res = queue[i].Val
			}
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return res
}
