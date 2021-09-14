package array

import "sort"

// 406
type ppl [][]int

func reconstructQueue(people [][]int) [][]int {
	p := ppl(people)
	sort.Sort(p)
	res := make([][]int, 0)
	for _, v := range p {
		newres := make([][]int, len(res[0:v[1]]))
		copy(newres, res[0:v[1]])
		newres = append(newres, v)
		newres = append(newres, res[v[1]:]...)
		res = newres
	}
	return res
}

func (p ppl) Len() int {
	return len(p)
}

func (p ppl) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ppl) Less(i, j int) bool {
	if p[i][0] > p[j][0] {
		return true
	} else if p[i][0] < p[j][0] {
		return false
	} else {
		return p[i][1] < p[j][1]
	}
}
