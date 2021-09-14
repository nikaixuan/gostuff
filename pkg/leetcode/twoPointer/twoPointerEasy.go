package twoPointer

import (
	"goproject/pkg/leetcode/linkedList"
)

// 26
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 1
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

// 83
func deleteDuplicates(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head.Next
	for right != nil {
		if left.Val != right.Val {
			left.Next = right
			left = left.Next

		}
		right = right.Next
	}
	left.Next = nil
	return head
}

// 27
func removeElement(nums []int, val int) int {
	left, right := -1, 0
	for right < len(nums) {
		if nums[right] != val {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

// 283
func moveZeroes(nums []int) {
	z, nz := 0, 0
	for nz < len(nums) {
		if nums[nz] == 0 {
			nz++
		}
		if nums[z] != 0 {
			z++
		}
		nums[nz], nums[z] = nums[z], nums[nz]
	}
}
