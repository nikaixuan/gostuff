package twoPointer

// 167
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	res := make([]int, 2)
	for i < j {
		if numbers[j]+numbers[i] > target {
			j--
		} else if numbers[j]+numbers[i] < target {
			i++
		} else if numbers[j]+numbers[i] == target {
			res[0] = i + 1
			res[1] = j + 1
			break
		}
	}
	return res
}
