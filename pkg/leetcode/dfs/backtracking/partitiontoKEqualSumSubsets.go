package backtracking

// 698
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	subsum := sum / k
	r := make([]int, k)
	return partitionBacktrack(nums, r, subsum, 0)
}

// traverse base on the number of the subset that need to split
func partitionBacktrack(nums, r []int, sum int, index int) bool {
	if index == len(nums) {
		for i := 0; i < len(r); i++ {
			if r[i] != sum {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(r); i++ {
		if r[i]+nums[index] > sum {
			continue
		}
		r[i] += nums[index]
		if partitionBacktrack(nums, r, sum, index+1) {
			return true
		}
		r[i] -= nums[index]
	}
	return false
}

// traverse base on the numbers
func canPartitionKSubsets2(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	subsum := sum / k
	r := make([]bool, len(nums))
	return partitionBacktrack2(nums, r, k, 0, 0, subsum)
}

// traverse base on the number of the subset that need to split
func partitionBacktrack2(nums []int, r []bool, k int, sum, index, target int) bool {
	if k == 0 {
		return true
	}
	if sum == target {
		return partitionBacktrack2(nums, r, k-1, 0, 0, target)
	}
	for i := index; i < len(nums); i++ {
		if r[i] || sum+nums[i] > target {
			continue
		}
		sum = sum + nums[i]
		r[i] = true
		// need the if return block so the program will not count [3,2,2] twice
		// once find a true answer, return to previous level
		if partitionBacktrack2(nums, r, k, sum, i+1, target) {
			return true
		}
		sum = sum - nums[i]
		r[i] = false
	}
	return false
}
