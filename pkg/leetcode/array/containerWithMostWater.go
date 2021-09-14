package array

// 11
func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		var h int
		if height[i] > height[j] {
			h = height[j]
		} else {
			h = height[i]
		}
		if h*(j-i) > max {
			max = h * (j - i)
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return max
}
