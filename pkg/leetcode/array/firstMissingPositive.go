package array

// 41
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 1; i <= n; i++ {
		if nums[i-1] > 0 && nums[i-1] <= n && nums[i-1] != nums[nums[i-1]-1] {
			nums[i-1], nums[nums[i-1]-1] = nums[nums[i-1]-1], nums[i-1]
			i--
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
