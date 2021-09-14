package strings

// 316
// stack
func removeDuplicateLetters(s string) string {
	// to count if there is any same string in the rest of s
	greedy := make([]int, 26)
	// to check whether the string is duplicate
	visited := make([]bool, 26)
	b := []byte(s)
	for _, v := range b {
		// prefill the count
		greedy[v-'a']++
	}
	stack := make([]byte, 0)
	for _, v := range b {
		index := v - 'a'
		greedy[index]--
		// if duplicate, skip, means this string is already in the result and seems valid for now
		if visited[index] {
			continue
		}

		// when the current string is smaller, and the previous string has another one to use later, pop the previous one
		for len(stack) > 0 && v < stack[len(stack)-1] && greedy[stack[len(stack)-1]-'a'] != 0 {
			// always use the large duplicate string as late as possible, so change the earlier one to false and wait for the later one
			visited[stack[len(stack)-1]-'a'] = false
			// pop from stack because we have another one to use later
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
		visited[index] = true
	}
	res := ""
	for i := 0; i < len(stack); i++ {
		res = res + string(stack[i])
	}
	return res
}

func removeDuplicateLetters2(s string) string {
	visited := make([]bool, 256)
	stack := make([]byte, 0)
	count := make([]int, 256)
	for _, v := range []byte(s) {
		count[v]++
	}
	for _, v := range []byte(s) {
		count[v]--
		if visited[v] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > v && count[stack[len(stack)-1]] > 0 {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
		visited[v] = true
	}
	res := ""
	for i := 0; i < len(stack); i++ {
		res = res + string(stack[i])
	}
	return res
}
