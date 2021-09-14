package linkedList

// 19
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// for edge case [1], 1
	fast, slow := head, head
	for i := n; i > 0; i-- {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
