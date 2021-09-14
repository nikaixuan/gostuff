package stack

// 402
// stack
// 单调栈
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	stack := make([]byte, 0)
	i := 0
	for i < len(num) {
		for len(stack) > 0 && k > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
		i++
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	res := ""
	for j := 0; j < len(stack); j++ {
		res = res + string(stack[j])
	}
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}

	return res
}
