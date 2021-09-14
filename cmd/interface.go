package main

import "fmt"

type myImplement struct{}

type flyer interface {
	fly()
}

type walker interface {
	walk()
}

type pig struct{}
type bird struct{}

func (p *pig) walk() {
	fmt.Println("pig walk")
}

func (b *bird) fly() {
	fmt.Println("bird walk")
}

func (m *myImplement) String() string {
	return "hi"
}

func getStringer() fmt.Stringer {
	var s *myImplement = nil
	if s == nil {
		fmt.Println("s is nil")
		// return nil
	} else {
		fmt.Println("s is not nil")
	}
	return s
}

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	}
}

func main() {
	if getStringer() == nil {
		fmt.Println("getStringer is nil")
	} else {
		fmt.Println("getStringer is not nil")
	}

	p1 := new(pig)
	var a walker
	a = p1
	p2 := a.(*pig)
	fmt.Println(p1, p2)
}
