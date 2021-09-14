package miscellaneous

import "fmt"

// 207
// topological sort
func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int)
	for i := range prerequisites {
		m[prerequisites[i][0]] = append(m[prerequisites[i][0]], prerequisites[i][1])
	}
	fmt.Println(m)
	for k, v := range m {
		for i := range v {
			if _, ok := m[v[i]]; ok {
				for j := range m[v[i]] {
					fmt.Println(k, m[v[i]][j])
					if k == m[v[i]][j] {
						return false
					}
				}
			}
		}
	}
	return true
}
