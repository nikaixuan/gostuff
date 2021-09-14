package datastructure

import "sync"

type Stack struct {
	array []string
	lock  sync.Mutex
	len   int
}

func (s *Stack) Push(a string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = append(s.array, a)

	s.len = len(s.array)
}

func (s *Stack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len < 1 {
		panic("stack is empty")
	}

	result := s.array[s.len-1]

	newArray := make([]string, s.len-1)
	// newArray = s.array[0:s.len-1]
	for i := 0; i < s.len-1; i++ {
		newArray[i] = s.array[i]
	}

	s.len = len(newArray)
	s.array = newArray
	return result
}

func (s *Stack) Peek() string {

	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == 0 {
		panic("empty")
	}

	return s.array[s.len]
}
