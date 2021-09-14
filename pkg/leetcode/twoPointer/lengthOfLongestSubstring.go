package twoPointer

import "math"

// 3
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	//res := 1
	//for i:=0;i< len(s);i++ {
	//	n := 0
	//	m := map[byte]struct{}{}
	//	for j:=i;j<len(s);j++ {
	//		if _, ok := m[s[j]]; !ok {
	//			m[s[j]] = struct{}{}
	//			n = n + 1
	//		} else {
	//			break
	//		}
	//	}
	//	res = int(math.Max(float64(res), float64(n)))
	//}

	res := 0
	m := make(map[byte]int)
	for i, j := 0, 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			j = int(math.Max(float64(j), float64(m[s[i]]+1)))
		}
		m[s[i]] = i
		res = int(math.Max(float64(res), float64(i-j+1)))
	}

	return res
}

func lengthOfLongestSubstring2(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		r := s[right]
		window[r]++
		right++
		// when repeat string appear, start shrink the window
		for window[r] > 1 {
			l := s[left]
			left++
			window[l]--
		}
		res = int(math.Max(float64(res), float64(right-left)))
	}
	return res
}
