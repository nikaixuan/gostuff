package divideAndConquer

import "strconv"

// 241
// map用于剪枝
var m = map[string][]int{}

func diffWaysToCompute(expression string) []int {
	// 查找结果，剪枝
	if _, ok := m[expression]; ok {
		return m[expression]
	}
	result := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			ls := diffWaysToCompute(expression[0:i])
			rs := diffWaysToCompute(expression[i+1:])
			for _, l := range ls {
				for _, r := range rs {
					res := 0
					switch expression[i] {
					case '+':
						res = l + r
					case '-':
						res = l - r
					case '*':
						res = l * r
					}
					result = append(result, res)
				}
			}
		}
	}
	// base case to end recursive
	// 当算式是一个数字没有符号时，返回此数字
	if len(result) == 0 {
		n, _ := strconv.Atoi(expression)
		result = []int{n}
	}
	// 记录结果
	m[expression] = result
	return result
}
