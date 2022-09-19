package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s, ok := i.(int)
	fmt.Println(s, ok)

	x := i.(string)
	fmt.Println(x)

}
