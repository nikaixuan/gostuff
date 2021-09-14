package binarySearch

// 153
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// although both left and right are close, not using <=, because not searching a target
	for left < right {
		mid := (left + right) / 2

		// mid cannot be the min value (at least large than nums[right] already)
		if nums[mid] > nums[right] {
			// squeeze right
			left = mid + 1
		} else {
			// nums[mid]<=nums[right] means mid might be the min or min is still on the left
			// cannot use right = mid - 1, that's also why cannot use left<=right in for loop
			// use left<=right and right = mid together will lead to infinite loop
			right = mid
		}
	}
	// the loop end when left=right, squeeze to min already
	return nums[left]
}
