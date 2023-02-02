package main

import "fmt"

func main() {
	var ch chan string

	ch <- "Hello"
	ch <- "World"

	data2, ok := <-ch
	fmt.Println("1:", data2, ok)
	data2, ok = <-ch
	fmt.Println("2:", data2, ok)
}
