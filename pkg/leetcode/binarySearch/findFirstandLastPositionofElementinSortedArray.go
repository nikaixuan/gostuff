package binarySearch

// 34
func searchRange(nums []int, target int) []int {
	n := len(nums)
	i, j := 0, n-1
	res := []int{-1, -1}
	if n == 0 {
		return res
	}
	for i < j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	if nums[i] != target {
		return res
	} else {
		res[0] = i
	}
	j = n - 1
	for i < j {
		mid := (i+j)/2 + 1
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid
		}
	}
	res[1] = j
	return res

}

func searchRangeBinarySearch(nums []int, target int) []int {
	n := len(nums)
	i, j := 0, n-1
	res := []int{-1, -1}
	if n == 0 {
		return res
	}
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		}
	}
	if i >= len(nums) || nums[i] != target {
		return res
	} else {
		res[0] = i
	}
	j = n - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		}
	}
	res[1] = j
	return res

}
