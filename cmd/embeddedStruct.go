package main

import "fmt"

type innerS struct {
	in1, in2 int
}

type outer struct {
	o1, o2 int
	int
	innerS
}

func main() {
	o := new(outer)
	o.o1 = 1
	o.o2 = 2
	o.int = 100
	fmt.Println(o)
}
