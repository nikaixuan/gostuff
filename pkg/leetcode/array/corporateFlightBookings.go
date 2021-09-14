package array

// 1109
// 差分数组
// array[i] = nums[i] - nums[i-1]
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	// 生成差分数组
	// 对nums[i, j]区间加减，可以通过在差分数组中array[i] += num and array[j+1] += -num表示
	for _, v := range bookings {
		res[v[0]-1] += v[2]
		if v[1] < n {
			res[v[1]] += -v[2]
		}
	}
	// 通过差分数组恢复
	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}
	return res
}
