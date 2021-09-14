package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is ", v.Kind() == reflect.Float64)
	fmt.Println("value is ", v.Float())
	y := float32(v.Float())
	fmt.Println(y)

	z := v.Interface().(float64)
	fmt.Println(z)
	// will throw exception, not accessible
	// v.SetFloat(3.41)

	vp := reflect.ValueOf(&x)
	vp.Elem().SetFloat(3.41)
	fmt.Println(x)
}
