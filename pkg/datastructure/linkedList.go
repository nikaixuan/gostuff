package datastructure

import "fmt"

type Ring struct {
	Val        int
	Next, Prev *Ring
}

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func New(n int) *Ring {
	if n < 1 {
		return nil
	}
	r := new(Ring)
	r.Val = 0
	p := r
	for i := 1; i < n; i++ {
		r.Next = &Ring{Prev: r, Val: i}
		r = r.Next
	}
	r.Next = p
	p.Prev = r
	return r
}

func (r *Ring) Print() {
	fmt.Println(r.Val)
	for p := *r.Next; p.Val != r.Val; p = *p.Next {
		fmt.Println(p.Val)
	}
}

func (r *Ring) Move(n int) *Ring {
	p := r
	if r.Next == nil {
		return nil
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			p = r.Prev
		}
	case n > 0:
		for ; n > 0; n-- {
			p = r.Next
		}
	}
	return p
}

func ReverseLinkListRec(head *LinkedList) *LinkedList {
	if head.Next == nil {
		return head
	}
	last := ReverseLinkListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseListIterative1(head *LinkedList) *LinkedList {
	var pre *LinkedList
	cur, nxt := head, head
	for nxt != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverseListIterative2(head *LinkedList) *LinkedList {
	var prev *LinkedList
	node := head

	for node != nil {
		temp := node.Next
		node.Next = prev

		prev = node
		node = temp
	}
	return prev
}
