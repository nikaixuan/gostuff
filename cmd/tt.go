package main

import (
	"fmt"
)

func main() {

	//a := &largeInt{3,5,6,2,34,2345,135,5123,23}
	//heap.Init(a)
	//heap.Push(a, 1)
	//fmt.Println(a)

	s := "abcd"
	fmt.Println(s[0])
	fmt.Println([]byte(s)[0])
	fmt.Println([]rune(s)[0])
	//m := make(map[byte][]int)
	//for i := range ring {
	//	if _, ok := m[key[i]]; !ok {
	//		m[key[i]] = []int{i}
	//	} else {
	//		curr := m[key[i]]
	//		m[key[i]] = append(curr, i)
	//	}
	//}

	//a := make([][]int, 0)
	//b := make([]int, 0)
	//for i := range []int{1,2,3,4,5} {
	//	// need to pass in the pointer
	//	b = append(b, i)
	//	temp := make([]int, 0)
	//	addarr(&a, append(temp, b...))
	//	b[0] = 100
	//}
	//fmt.Println(a)

}

func addarr(a *[][]int, b []int) {
	*a = append(*a, b)
}
