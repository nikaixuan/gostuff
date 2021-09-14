package twoPointer

// 567
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, v := range []byte(s1) {
		need[v]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		r := []byte(s2)[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}

		// maintain the window length is s1 length
		for right-left >= len(s1) {
			l := []byte(s2)[left]
			if valid == len(need) {
				return true
			}
			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
			left++
		}
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	v1 := make([]int, 26)
	v2 := make([]int, 26)
	for i := range s1 {
		v1[s1[i]-'a']++
		v2[s2[i]-'a']++
	}
	if match(v1, v2) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		// sliding window, index move right
		v2[s2[i]-'a']++
		// left most element delete
		v2[s2[i-len(s1)]-'a']--
		if match(v2, v1) {
			return true
		}
	}
	return false
}

func match(v1, v2 []int) bool {
	for i := 0; i < 26; i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
