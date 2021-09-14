package designDataStructure

type monoQueue struct {
	queue []int
}

// 239
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	queue := &monoQueue{
		queue: make([]int, 0),
	}
	for i := 0; i < len(nums); i++ {
		for i < k-1 {
			queue.push(nums[i])
			i++
		}
		queue.push(nums[i])
		res = append(res, queue.max())
		queue.pop(nums[i-k+1])
	}
	return res
}

func (m monoQueue) max() int {
	return m.queue[0]
}

// when push a number in the queue, if it is max number, make sure it squeeze all smaller number
// and become the 0 index
func (m *monoQueue) push(n int) {
	for len(m.queue) > 0 && m.queue[len(m.queue)-1] < n {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, n)
}

func (m *monoQueue) pop(n int) {
	if m.queue[0] == n {
		m.queue = m.queue[1:]
	}
}
