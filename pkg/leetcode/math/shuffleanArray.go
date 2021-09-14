package math

import "math/rand"

// 384
type SolutionShuffle struct {
	origin []int
}

func ConstructorShuffle(nums []int) SolutionShuffle {
	return SolutionShuffle{
		origin: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *SolutionShuffle) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *SolutionShuffle) Shuffle() []int {
	res := make([]int, len(this.origin))
	temp := append([]int{}, this.origin...)
	for i := len(this.origin); i > 0; i-- {
		// gen random index
		idex := rand.Intn(i)
		// swap that index number with last index number
		temp[len(temp)-1], temp[idex] = temp[idex], temp[len(temp)-1]
		// pop that number and add to array
		res[i-1] = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
	}
	return res
}

func (this *SolutionShuffle) Shuffle2() []int {
	temp := make([]int, len(this.origin))
	copy(temp, this.origin)
	for i := 0; i < len(this.origin)-1; i++ {
		// gen random index
		idex := rand.Intn(len(this.origin) - i)
		// swap that index number with last index number
		temp[len(temp)-1-i], temp[idex] = temp[idex], temp[len(temp)-1-i]
	}
	return temp
}
