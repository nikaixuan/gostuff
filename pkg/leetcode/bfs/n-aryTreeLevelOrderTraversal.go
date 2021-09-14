package bfs

// 429
type naryNode struct {
	Val      int
	Children []*naryNode
}

func levelOrder429(root *naryNode) [][]int {
	queue := make([]*naryNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	for len(queue) > 0 {
		l := len(queue)
		r := make([]int, 0)
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			r = append(r, curr.Val)
			for _, node := range curr.Children {
				if node != nil {
					queue = append(queue, node)
				}
			}
		}
		res = append(res, r)
	}
	return res
}
