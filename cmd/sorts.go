package main

import (
	"fmt"
	"sort"
)

type carKind int

const (
	BMW = iota
	VM
	TOYOTA
)

type car struct {
	kind carKind
	name string
}

func main() {
	a := sort.IntSlice{3, 5, 1, 4, 2}
	a.Sort()
	fmt.Println(a)

	cars := []car{
		{VM, "golf"},
		{BMW, "m3"},
		{TOYOTA, "camery"},
	}
	sort.Slice(cars, func(i, j int) bool {
		if cars[i].kind != cars[j].kind {
			return cars[i].kind < cars[j].kind
		}
		return cars[i].name < cars[j].name
	})
	fmt.Println(cars)
}
