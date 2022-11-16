package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyType struct {
	Name string
}

func (MyType) Method() {}

type User struct{}

func (User) Method() {}

func main() {
	t := MyType{}
	var i MyInterface = MyType{}

	fmt.Println(t == i)

	var j MyInterface = MyType{"HH"}

	fmt.Println(t == j)

	var k MyInterface = User{}

	fmt.Println(t == k)
}
