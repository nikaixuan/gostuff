package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

// 不用每次进位，只需记录/10的得数在和后面两个数加起来，到最后再进位
func add(l1 *ListNode, l2 *ListNode, v int) *ListNode {
	if l1 == nil && l2 == nil {
		if v == 1 {
			return &ListNode{
				Val:  1,
				Next: nil,
			}
		} else {
			return nil
		}
	}
	if l1 == nil {
		newv := 0
		if l2.Val+v > 9 {
			newv = 1
		}
		return &ListNode{
			Val:  l2.Val + v,
			Next: add(nil, l2.Next, newv),
		}
	}
	if l2 == nil {
		newv := 0
		if l1.Val+v > 9 {
			newv = 1
		}
		return &ListNode{
			Val:  l1.Val + v,
			Next: add(l1.Next, nil, newv),
		}
	}
	newv := 0
	if l1.Val+l2.Val+v > 9 {
		newv = 1
	}
	return &ListNode{
		Val:  l1.Val + +l2.Val + v,
		Next: add(l1.Next, l2.Next, newv),
	}
}
