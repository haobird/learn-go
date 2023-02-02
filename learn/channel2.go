package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	close(ch)

	ch <- "Three"

	data2, ok := <-ch
	fmt.Println("1:", data2, ok)
	data2, ok = <-ch
	fmt.Println("2:", data2, ok)

	data2, ok = <-ch
	fmt.Println("3:", data2, ok)

	data2, ok = <-ch
	fmt.Println("4:", data2, ok)

	// for data := range ch {
	// 	fmt.Println(data)
	// }
}
