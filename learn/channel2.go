package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	close(ch)

	for data := range ch {
		fmt.Println(data)
	}
}
