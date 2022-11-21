package main

import "fmt"

type User struct {
	Name string
}

func main() {
	var drivers map[string]User
	fmt.Println(drivers, drivers == nil)
	fmt.Printf("%v\n", drivers)
	drivers = make(map[string]User)
	fmt.Println(drivers, drivers == nil)
	fmt.Printf("%v\n", drivers)
	drivers1 := new(map[string]User)
	fmt.Println(drivers1, drivers1 == nil)
	fmt.Printf("%v\n", drivers1)

	var m = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m["c"] = User{"Steve"}
	m["d"] = User{Name: "xxx"}

	fmt.Println(m)

	c, ok := m["c"]
	fmt.Println("Key c:", c, ok)

	d, ok := m["d"]
	fmt.Println("Key d:", d, ok)

	var m1 = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	m2 := m1
	m2["c"] = User{"Steve"}

	fmt.Println(m1) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
	fmt.Println(m2) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
}
