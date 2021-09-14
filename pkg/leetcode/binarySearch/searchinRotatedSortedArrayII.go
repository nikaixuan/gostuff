package binarySearch

// 81
func search2(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		//If we know for sure right side is sorted or left side is unsorted
		if nums[mid] < nums[right] || nums[mid] < nums[left] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
			//If we know for sure left side is sorted or right side is unsorted
		} else if nums[mid] > nums[left] || nums[mid] > nums[right] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			//If we get here, that means nums[start] == nums[mid] == nums[end], then shifting out
			//any of the two sides won't change the result but can help remove duplicate from
			//consideration, here we just use end-- but left++ works too
		} else {
			left++
		}
	}
	return false
}
