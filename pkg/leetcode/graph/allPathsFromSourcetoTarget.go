package graph

// 797
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	r := make([]int, 0)
	traverseGraph(graph, 0, r, &res)
	return res
}

func traverseGraph(graph [][]int, curr int, r []int, res *[][]int) {
	n := len(graph) - 1

	r = append(r, curr)
	if curr == n {
		temp := make([]int, len(r))
		copy(temp, r)
		*res = append(*res, temp)
		return
	}

	for _, node := range graph[curr] {
		traverseGraph(graph, node, r, res)
	}
}
