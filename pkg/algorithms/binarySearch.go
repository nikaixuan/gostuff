package algorithms

func BinarySearchBothClose(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// terminate when left > right, search array is empty
	for left <= right {
		mid := left + (right-left)/2

		if target < nums[mid] {
			// mid is checked, so skip mid
			right = mid - 1
		} else if target > nums[mid] {
			// mid is checked, so skip mid
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchLeftBound(nums []int, target int) int {
	// right bound is open
	left, right := 0, len(nums)

	// terminate when left == right, close on left
	for left < right {
		mid := left + (right-left)/2
		// left is important here, means how many number in the array is larger than the target, could be from 0 to len(nums)
		if nums[mid] == target {
			// because it is open on right, so no need to -1
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func BinarySearchLeftBoundAllClose(nums []int, target int) int {
	// right bound is close
	left, right := 0, len(nums)-1

	// terminate when left > right, max left can be len(nums)
	for left <= right {
		mid := left + (right-left)/2
		// left is important here, means how many number in the array is larger than the target, could be from 0 to len(nums)
		if nums[mid] == target {
			// because it is close on right, so need to -1
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func BinarySearchRightBound(nums []int, target int) int {
	// right bound is open
	left, right := 0, len(nums)

	// terminate when left == right, close on left
	for left < right {
		mid := left + (right-left)/2
		// left is important here, it will always search to right to find the biggest index that match the target
		if nums[mid] == target {
			// because it is open on right, so left needs to +1
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left-1 < 0 || nums[left-1] != target {
		return -1
	}
	return left - 1
}

func BinarySearchRightBoundAllClose(nums []int, target int) int {
	// right bound is open
	left, right := 0, len(nums)-1

	// terminate when left == right, close on left
	for left <= right {
		mid := left + (right-left)/2
		// left is important here, it will always search to right to find the biggest index that match the target
		if nums[mid] == target {
			// because it is open on right, so left needs to +1
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left-1 < 0 || nums[left-1] != target {
		return -1
	}
	return left - 1
}
