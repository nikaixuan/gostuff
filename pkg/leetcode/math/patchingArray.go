package math

// 330
// [1,2,4,13,43] n=100
// [1,2,4] cover 7, add 8 so cover [1...16)
// [1,2,4,(8),13] cover [1...16+13), then add 29 to cover [1...29+29)
// cuz [1...29+29+43) include 100, so finish
func minPatches(nums []int, n int) int {
	sum, add, i := 1, 0, 0
	for sum <= n {
		if i < len(nums) && nums[i] <= sum {
			sum += nums[i]
			i++
		} else {
			sum += sum
			add++
		}
	}
	return add
}
