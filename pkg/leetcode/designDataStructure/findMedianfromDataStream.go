package designDataStructure

import (
	"container/heap"
)

// 295
type largeInt []int

func (l *largeInt) Len() int {
	return len(*l)
}

func (l *largeInt) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *largeInt) Push(x interface{}) {
	*l = append(*l, x.(int))
}

func (l *largeInt) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *largeInt) Less(i, j int) bool {
	return (*l)[j] > (*l)[i]
}

type smallInt []int

func (s *smallInt) Less(i, j int) bool {
	return (*s)[i] > (*s)[j]
}

func (s *smallInt) Len() int {
	return len(*s)
}

func (s *smallInt) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *smallInt) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *smallInt) Pop() interface{} {
	n := len(*s)
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return x
}

type MedianFinder struct {
	Large *largeInt
	Small *smallInt
}

/** initialize your data structure here. */
func ConstructorMedianFinder() MedianFinder {
	m := MedianFinder{}
	m.Small = (*smallInt)(new([]int))
	m.Large = (*largeInt)(new([]int))
	return m
}

func (this *MedianFinder) AddNum(num int) {
	if this.Small.Len() >= this.Large.Len() {
		heap.Push(this.Small, num)
		heap.Push(this.Large, heap.Pop(this.Small))
	} else {
		heap.Push(this.Large, num)
		heap.Push(this.Small, heap.Pop(this.Large))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Small.Len() > this.Large.Len() {
		return float64((*this.Small)[0])
	} else if this.Small.Len() < this.Large.Len() {
		return float64((*this.Large)[0])
	}
	return (float64((*this.Small)[0]) + float64((*this.Large)[0])) / 2
}
