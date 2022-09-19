package main

import "fmt"

func main() {
	WillPanic()
}

func handlePanic() {
	data := recover()
	fmt.Println("Recovered:", data)
}

func WillPanic() {
	defer handlePanic()

	panic("Woah")
}
