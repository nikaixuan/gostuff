package backtracking

import "sort"

// 332
func findItinerary(tickets [][]string) []string {
	m1 := make(map[string][]string)
	m2 := make(map[string][]bool)
	res := make([]string, 0)
	n := len(tickets)
	for _, v := range tickets {
		m1[v[0]] = append(m1[v[0]], v[1])
		m2[v[0]] = append(m2[v[0]], false)
	}
	for _, v := range m1 {
		sort.Strings(v)
	}
	curr := "JFK"
	res = append(res, curr)
	itinerBackTrack(m1, m2, &res, n, curr)
	return res
}

func itinerBackTrack(m1 map[string][]string, m2 map[string][]bool, res *[]string, n int, curr string) bool {
	if n == 0 {
		return true
	}
	if len(m1[curr]) == 0 {
		return false
	}
	for i, v := range m2[curr] {
		if !v {
			m2[curr][i] = true
			*res = append(*res, m1[curr][i])
			// 因为是排过序的，所以第一个返回的结果一定是lex order最小的
			if itinerBackTrack(m1, m2, res, n-1, m1[curr][i]) {
				return true
			}
			m2[curr][i] = false
			*res = (*res)[:len(*res)-1]
		}
	}
	return false
}
