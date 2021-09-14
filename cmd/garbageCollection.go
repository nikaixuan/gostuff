package main

import (
	"fmt"
	"runtime"
)

type road int

func findRoad(r *road) {
	fmt.Println("road: ", *r)
}

func entry() {
	var rd = road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
}

func main() {
	entry()
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
}
