package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func (i *MyInt) Update(value int) {
	*i = 15
}

func main() {
	i := MyInt(10)

	fmt.Println(i.isGreater(11))

	i.Update(10)

	fmt.Println(i)
	fmt.Println(i.isGreater(11))
}
