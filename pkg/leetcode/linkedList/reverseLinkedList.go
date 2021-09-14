package linkedList

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first := &ListNode{Val: head.Val}
	return goReverse(first, head.Next)
}

func goReverse(curr *ListNode, old *ListNode) *ListNode {
	if old == nil {
		return curr
	}
	return goReverse(&ListNode{Val: old.Val, Next: curr}, old.Next)
}
