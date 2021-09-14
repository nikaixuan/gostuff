package designDataStructure

import "goproject/pkg/algorithms"

// 990
// union find
func equationsPossible(equations []string) bool {
	u := algorithms.Uf{}
	u.Uf(26)
	for _, s := range equations {
		b := []byte(s)
		if b[1] == '=' {
			u.Union(int(b[0]-'a'), int(b[3]-'a'))
		}
	}
	for _, s := range equations {
		b := []byte(s)
		if b[1] == '!' {
			if u.Connected(int(b[0]-'a'), int(b[3]-'a')) {
				return false
			}
		}
	}
	return true
}
