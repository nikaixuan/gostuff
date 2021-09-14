package linkedList

// 234
var left *ListNode

// 后序遍历法
func isPalindromeLinkList1(head *ListNode) bool {
	left = head
	return paliPostorder(head)
}

func paliPostorder(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := paliPostorder(right.Next)
	res = res && left.Val == right.Val
	left = left.Next
	return res
}

// 快慢指针找中心，反转链表并比较法
func isPalindromeLinkList2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left, right := head, reverseLL(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseLL(head *ListNode) *ListNode {
	node := head
	var pre *ListNode
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
