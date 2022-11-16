package main

import (
	"fmt"
	"reflect"
)

type MyInt int

type Person struct {
	Name string
}

var (
	x        = byte('x')
	y uint8  = 'x'
	z MyInt  = 7
	p Person = Person{Name: "dd"}
)

func main() {
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind:", v.Kind())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	// x = uint8(v.Int())
	v1 := reflect.ValueOf(z)
	fmt.Println("type:", v1.Type())                        // int.
	fmt.Println("kind:", v1.Kind())                        // int.
	fmt.Println("kind is int: ", v1.Kind() == reflect.Int) // true.

	v2 := reflect.ValueOf(p)
	fmt.Println("type:", v2.Type()) // int.
	fmt.Println("kind:", v2.Kind()) // int.
	// fmt.Println("kind is int: ", v1.Kind() == reflect.Int) // true.
}
