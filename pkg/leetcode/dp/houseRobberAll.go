package dp

import (
	"goproject/pkg/leetcode/binaryTree"
	"math"
)

// 198
func rob1(nums []int) int {

	// standard dp
	//dp := make([]int, len(nums))
	//dp[0] = nums[0]
	//for i:=1;i<len(nums);i++ {
	//	if i==1 {
	//		dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i])))
	//		continue
	//	}
	//	dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i]+dp[i-2])))
	//}
	//return dp[len(nums)-1]

	// less space solution
	dp0, dp1, dp2 := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		dp0 = int(math.Max(float64(dp1), float64(dp2+nums[i])))
		dp2 = dp1
		dp1 = dp0
	}
	return dp0
}

// 213
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// compare rob first not last and rob last but not first
	res := int(math.Max(float64(rob1(nums[0:len(nums)-1])), float64(rob1(nums[1:]))))
	return res
}

// 337
func rob3(root *binaryTree.TreeNode) int {
	res := robDfs(root)
	return int(math.Max(float64(res[0]), float64(res[1])))
}

func robDfs(root *binaryTree.TreeNode) []int {
	if root == nil {
		return make([]int, 2)
	}
	left := robDfs(root.Left)
	right := robDfs(root.Right)
	res := make([]int, 2)
	res[0] = int(math.Max(float64(left[0]), float64(left[1])) + math.Max(float64(right[0]), float64(right[1])))
	res[1] = left[0] + right[0] + root.Val
	return res
}
