package bfs

// 752
// BFS
func openLock(deadends []string, target string) int {
	queue := make([]string, 0)
	n := -1
	m := make(map[string]struct{})
	queue = append(queue, "0000")
	for _, v := range deadends {
		m[v] = struct{}{}
	}
	for len(queue) != 0 {
		l := len(queue)
		n = n + 1
		for i := 0; i < l; i++ {
			head := queue[0]
			queue = queue[1:]
			if _, ok := m[head]; ok {
				continue
			}
			m[head] = struct{}{}
			if head == target {
				return n
			}
			for a := 0; a < len(head); a++ {
				temp1 := []rune(head)
				temp1[a] = rune((head[a]-'0'+1+10)%10 + '0')
				queue = append(queue, string(temp1))
				temp2 := []rune(head)
				temp2[a] = rune((head[a]-'0'-1+10)%10 + '0')
				queue = append(queue, string(temp2))
			}
		}
	}
	return -1
}
