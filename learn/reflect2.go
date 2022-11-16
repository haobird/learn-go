package main

import (
	"fmt"
	"reflect"
)

type MyInt int

var (
	a MyInt = 10
)

func main() {
	fmt.Println("type kind", reflect.TypeOf(a).Kind(), "value kind", reflect.ValueOf(a).Kind())
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	fmt.Println("type:", t, "kind", t.Kind(), "size", t.Size())

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type(), "kind", v.Kind(), "value", v.String())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
