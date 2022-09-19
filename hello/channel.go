package main

import "fmt"

func speak(arg string, ch chan<- string) {
	ch <- arg // Send Only
}

func main() {
	ch := make(chan string, 2)

	// go func() {
	// 	data3, ok := <-ch
	// 	fmt.Println(data3, ok)
	// }()

	go speak("Hello World", ch)
	go speak("Hi again", ch)

	data1 := <-ch
	fmt.Println(data1)

	data2, ok := <-ch
	fmt.Println(data2, ok)

	fmt.Println(ch)

	close(ch)

	data3, ok := <-ch
	fmt.Println(data3, ok)
	fmt.Println(ch)

}
