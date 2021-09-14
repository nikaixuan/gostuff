package designDataStructure

import (
	"math"
)

// 895
type FreqStack struct {
	maxFreq   int
	valToFreq map[int]int
	freqToVal map[int]*[]int
}

func ConstructorFreqStack() FreqStack {
	fs := FreqStack{
		maxFreq:   0,
		valToFreq: make(map[int]int),
		freqToVal: make(map[int]*[]int),
	}
	return fs
}

func (this *FreqStack) Push(val int) {
	freq, ok := this.valToFreq[val]
	if !ok {
		freq = 0
	}
	freq = freq + 1
	this.valToFreq[val] = freq
	if _, ok := this.freqToVal[freq]; !ok {
		this.freqToVal[freq] = new([]int)
	}
	*this.freqToVal[freq] = append(*this.freqToVal[freq], val)
	this.maxFreq = int(math.Max(float64(this.maxFreq), float64(freq)))
}

func (this *FreqStack) Pop() int {
	vals := *this.freqToVal[this.maxFreq]
	val := vals[len(vals)-1]
	*this.freqToVal[this.maxFreq] = vals[:len(vals)-1]
	freq := this.valToFreq[val] - 1
	this.valToFreq[val] = freq
	if len(*this.freqToVal[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	return val
}
