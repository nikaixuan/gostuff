package algorithms

import "math"

func slidingWindow(s string, t string) string {
	// current string-count map inside the window
	window := make(map[byte]int)

	// string-count map is needed, based on the t string, should not be changed
	need := make(map[byte]int)
	for _, b := range []byte(t) {
		need[b]++
	}
	// left close, right open, valid is to compare how many elem in window is equal to need, which is valid.
	left, right, valid := 0, 0, 0
	// start index and the length of the possible result
	start, resl := 0, math.MaxInt64
	// stop when right hit the end
	for right < len(s) {
		r := []byte(s)[right]
		// right add first so it is open on right: [left, right)
		right++

		// check and update current window to the right
		// if current element is needed
		if _, ok := need[r]; ok {
			// update window map
			window[r]++
			// if the element number is equal, it is valid
			if window[r] == need[r] {
				valid++
			}
		}

		// shrink the window from the left
		// check whether the window needs shrink by checking the valid element number
		// if not enough, will keep expand right to search
		for valid == len(need) {
			// current left element
			l := []byte(s)[left]
			//compare the length to find shortest
			if right-left < resl {
				// if the current string length is shortest for now, change the start index
				start = left
				// update the shortest length
				resl = right - left
			}
			// shrink the window from left to find the shortest
			if _, ok := need[l]; ok {
				// if the shrink affect any element, we need to make it not valid FIRST!!
				// !! because we will delete an element anyway, so if the current element is just needed, the window is not valid anymore!!
				if window[l] == need[l] {
					valid--
				}
				// update element in window
				window[l]--
			}
			// update left index
			left++
		}
	}
	// if length is never updated
	if resl == math.MaxInt64 {
		return ""
	}
	// left close, right open
	return s[start : start+resl]
}
