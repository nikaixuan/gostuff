package array

// 75
func sortColors(nums []int) {
	z, o, t := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			z, o, t = z+1, o+1, t+1
			nums[t] = 2
			nums[o] = 1
			nums[z] = 0
		} else if nums[i] == 1 {
			o, t = o+1, t+1
			nums[t] = 2
			nums[o] = 1
		} else if nums[i] == 2 {
			t = t + 1
			nums[t] = 2
		}
	}
}
