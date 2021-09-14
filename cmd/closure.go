package main

import "fmt"

func accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {

	// closure example
	str := "abc"

	foo := func() {
		str = "bcd"
	}
	fmt.Println(str)
	foo()
	fmt.Println(str)

	// accumulator example
	accumulator := accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	accumulator2 := accumulate(2)
	fmt.Println(accumulator2())
}
