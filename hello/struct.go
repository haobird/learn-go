package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

type Person2 struct {
	FirstName, LastName string
	Age                 int
}

type SuperHero struct {
	Person
	Alias string
}

type superHero struct {
	Person Person
	Alias  string
}

func main() {
	p := Person{"Bruce", "Wayne", 40}
	p2 := Person2{"Bruce", "Wayne", 40}

	fmt.Println(p == p2)

	// s := SuperHero{}

	// s.FirstName = "Bruce"
	// s.LastName = "Wayne"
	// s.Age = 40
	// s.Alias = "batman"

	// fmt.Println(s)
}
