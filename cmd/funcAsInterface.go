package main

import "fmt"

type invoker interface {
	Call(interface{})
}

type testStruct struct {
}

func (t *testStruct) Call(i interface{}) {
	fmt.Println("call struct ", i)
}

// for function type, no need to initialise, just transfer other function(normal, anonymous, closure) to this func type
type funcCaller func(interface{})

func (f funcCaller) Call(i interface{}) {
	f(i)
}

func main() {
	var invo invoker

	invo = new(testStruct)
	invo.Call(1)

	invo = funcCaller(func(i interface{}) {
		fmt.Println("call interface ", i)
	})

	invo.Call(2)

}
