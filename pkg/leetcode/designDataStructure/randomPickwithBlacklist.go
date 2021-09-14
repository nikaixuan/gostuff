package designDataStructure

import "math/rand"

type Solution struct {
	sz int
	m  map[int]int
}

// 710
func ConstructorRandPick(n int, blacklist []int) Solution {
	size := n - len(blacklist)
	s := Solution{
		sz: size,
		m:  make(map[int]int, len(blacklist)),
	}
	last := n - 1
	for i := range blacklist {
		// map of blacklist number to index
		s.m[blacklist[i]] = 0
	}
	for _, v := range blacklist {
		if v >= size {
			continue
		}
		// map right blacklist number to first left not-blacklist number
		for _, ok := s.m[last]; ok; {
			last = last - 1
			_, ok = s.m[last]
		}
		// map one of the blacklist number to this index
		s.m[v] = last
		last = last - 1
	}
	return s
}

func (this *Solution) Pick() int {
	n := rand.Intn(this.sz)
	if _, ok := this.m[n]; ok {
		return this.m[n]
	}
	return n
}
