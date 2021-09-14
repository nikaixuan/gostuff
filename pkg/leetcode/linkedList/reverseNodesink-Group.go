package linkedList

// 25
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := k; i > 0; i-- {
		if b == nil {
			return head
		}
		b = b.Next
	}
	curr := reverseListab(a, b)
	a.Next = reverseKGroup(b, k)
	return curr
}

// iter
func reverseListab(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur, nxt := a, a
	for nxt != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// rec
func reverseListabrec(a, b *ListNode) *ListNode {
	if a.Next == b {
		return a
	}
	last := reverseListabrec(a.Next, b)
	a.Next.Next = a
	a.Next = b
	return last
}
