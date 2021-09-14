package binaryTree

// 116
type NextNode struct {
	Val   int
	Left  *NextNode
	Right *NextNode
	Next  *NextNode
}

// Level traversal
func connect1(root *NextNode) *NextNode {
	connectDfs(root, nil)
	return root
}

// pre-order
func connectDfs(c *NextNode, n *NextNode) {
	if c == nil {
		return
	}
	c.Next = n
	var nn *NextNode
	if c.Next != nil {
		nn = c.Next.Left
	} else {
		nn = nil
	}
	connectDfs(c.Left, c.Right)
	connectDfs(c.Right, nn)

}

// bfs
func connect2(root *NextNode) *NextNode {
	queue := make([]*NextNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		s := len(queue)
		for i := 0; i < s; i++ {
			n := queue[i]
			if i == s-1 {
				n.Next = nil
				queue = queue[i+1:]
			} else if i < s-1 {
				n.Next = queue[i+1]
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}
	return root
}
