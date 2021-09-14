package graph

// 310
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	m := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		m[i] = make(map[int]struct{})
	}
	for _, e := range edges {
		m[e[0]][e[1]] = struct{}{}
		m[e[1]][e[0]] = struct{}{}
	}
	leaves := make([]int, 0)
	for i, v := range m {
		if len(v) == 1 {
			leaves = append(leaves, i)
		}
	}
	for n > 2 {
		newLeaves := make([]int, 0)
		for _, l := range leaves {
			for k, _ := range m[l] {
				delete(m[k], l)
				delete(m[l], k)
				if len(m[k]) == 1 {
					newLeaves = append(newLeaves, k)
				}
			}
			n = n - 1
		}
		leaves = newLeaves
	}
	return leaves
}
