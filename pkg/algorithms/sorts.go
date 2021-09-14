package algorithms

// o(nlogn)
func quickSort(nums []int) {
	quick(nums, 0, len(nums)-1)
}

// 类似pre-order traversal
func quick(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	quick(nums, p+1, hi)
	quick(nums, lo, p-1)
}

func partition(nums []int, lo, hi int) int {
	// 不检查的话后面for循环需要检查索引
	if lo == hi {
		return lo
	}
	// 因为先加减，所以索引越界
	i, j := lo, hi+1
	pivot := nums[lo]
	for {
		// 从左开始查找pivot的位置，通过for循环找到比pivot大的
		i++
		for nums[i] < pivot {
			if i == hi {
				break
			}
			i++
		}
		// 从右开始查找pivot的位置，通过for循环找到比pivot小的
		j--
		for nums[j] > pivot {
			if j == lo {
				break
			}
			j--
		}
		// 如果有等于pivot的元素会出现i==j，否则会有i>j
		// 两种情况都应打破循环，因为此时j的左边（不包含j）全部小于pivot而右边全部大于pivot
		if i >= j {
			break
		}
		// 如果走到这里，一定有：
		// nums[i] > pivot && nums[j] < pivot 且i在左j在右
		// 所以需要交换 nums[i] 和 nums[j]，
		// 保证 nums[lo..i] < pivot < nums[j..hi]
		nums[i], nums[j] = nums[j], nums[i]
	}
	// partition完成，交换lo和j的元素，此时j左边全部小于等于pivot，右边大于pivot
	// 返回j，可以对j两边进行排序
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
