package bfs

// 279
func numSquares(n int) int {
	res := 0
	queue := make([]int, 0)
	m := make(map[int]struct{})
	queue = append(queue, 0)
	for len(queue) > 0 {
		l := len(queue)
		res = res + 1
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			j := 1
			for curr+j*j <= n {
				temp := curr + j*j
				if temp == n {
					return res
				}
				j = j + 1
				m[temp] = struct{}{}
				queue = append(queue, temp)
			}
		}
	}
	return -1
}
