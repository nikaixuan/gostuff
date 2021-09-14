package twoPointer

// 438
func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}
	res := make([]int, 0)
	left, right, valid := 0, 0, 0
	for right < len(s) {
		r := s[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}
		for right-left >= len(p) {
			l := s[left]
			if len(need) == valid {
				res = append(res, left)
			}
			if _, ok := window[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
			left++
		}
	}
	return res
}

func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	v1 := make([]int, 26)
	v2 := make([]int, 26)
	for i := range p {
		v1[p[i]-'a']++
		v2[s[i]-'a']++
	}
	res := make([]int, 0)
	if match2(v1, v2) {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		v2[s[i]-'a']++
		v2[s[i-len(p)]-'a']--
		if match2(v2, v1) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func match2(v1, v2 []int) bool {
	for i := 0; i < 26; i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
