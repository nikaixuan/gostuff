package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	newMap := make(map[*Node]*Node)
	newNode := &Node{Val: node.Val, Neighbors: []*Node{}}
	newMap[node] = newNode
	cloneDfs(node, newMap)
	return newNode
}

func cloneDfs(node *Node, nodeMap map[*Node]*Node) {
	if node == nil {
		return
	}
	for i := range node.Neighbors {
		if _, ok := nodeMap[node.Neighbors[i]]; !ok {
			n := &Node{Val: node.Neighbors[i].Val, Neighbors: []*Node{}}
			nodeMap[node.Neighbors[i]] = n
			cloneDfs(node.Neighbors[i], nodeMap)
		}
		nodeMap[node].Neighbors = append(nodeMap[node].Neighbors, nodeMap[node.Neighbors[i]])
	}
}
