package designDataStructure

import "math/rand"

// 380
type RandomizedSet struct {
	nums    []int
	sizeMap map[int]int
}

/** Initialize your data structure here. */
func ConstructorRandomizedSet() RandomizedSet {
	r := RandomizedSet{
		nums:    make([]int, 0),
		sizeMap: make(map[int]int),
	}
	return r
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.sizeMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.sizeMap[val] = len(this.nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.sizeMap[val]; !ok {
		return false
	}
	i := this.sizeMap[val]
	this.nums[i], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[i]
	this.sizeMap[this.nums[i]] = i
	delete(this.sizeMap, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.nums))
	return this.nums[n]
}
