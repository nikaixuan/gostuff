package strings

// 49
func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, v := range strs {
		k := [26]int{}
		for i := 0; i < len(v); i++ {
			k[v[i]-'a'] += 1
		}
		m[k] = append(m[k], v)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
