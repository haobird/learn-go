package main

import "fmt"

func main() {
	fmt.Println("dd")
	one := make(chan string)
	two := make(chan string)
	select {}

	close(one)
	close(two)
}
