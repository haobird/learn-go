package main

import "fmt"

type Car struct {
	Name string
	Year int
}

func (c Car) IsLatest() bool {
	return c.Year >= 2017
}

func (c *Car) UpdateName(name string) {
	c.Name = name
}

func main() {
	c := Car{"Tesla", 2021}

	c.UpdateName("Toyota")
	fmt.Println("Car:", c)

	c2 := &Car{"111", 2021}
	c2.UpdateName("222")
	fmt.Println("Car:", c2)

	fmt.Println("IsLatest", c2.IsLatest())
}
