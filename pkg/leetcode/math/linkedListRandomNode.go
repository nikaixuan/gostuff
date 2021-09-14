package math

import (
	"goproject/pkg/leetcode/linkedList"
	"math/rand"
)

// 382
// resovior sampling 水塘抽样
type Solution struct {
	head *linkedList.ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *linkedList.ListNode) Solution {
	return Solution{
		head: head,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	len, res := 0, 0
	h := this.head
	for h != nil {
		len++
		// generate a int between [0, n)
		// chance of it is 1/n
		if rand.Intn(len) == 0 {
			res = h.Val
		}
		h = h.Next
	}
	return res
}
