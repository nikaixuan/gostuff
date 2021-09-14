package array

// 560
// 前缀和
func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	preSum := make(map[int]int)
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		oldSum := sum - k
		if _, ok := preSum[oldSum]; ok {
			res += preSum[oldSum]
		}
		preSum[sum] += 1
	}
	return res
}
