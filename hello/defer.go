package main

import "fmt"

func main() {
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you?")

	fmt.Println("Doing some work...")
}
