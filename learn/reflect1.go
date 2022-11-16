package main

import (
	"fmt"
	"reflect"
)

var (
	x float64 = 3.4
	y string  = "xx"
	z int64   = 5
)

func main() {
	v := reflect.ValueOf(z)
	fmt.Printf("value is %7.1e\n", v.Interface())
}
