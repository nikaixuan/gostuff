package math

import "math/rand"

// 398
// 水塘抽样
type SolutionRandomPick struct {
	curr []int
}

func ConstructorRandomPick(nums []int) SolutionRandomPick {
	return SolutionRandomPick{
		curr: nums,
	}
}

func (this *SolutionRandomPick) Pick(target int) int {
	seed, res := 0, 0
	for i, v := range this.curr {
		if v == target {
			seed++
			if rand.Intn(seed) == 0 {
				res = i
			}
		}
	}
	return res
}
