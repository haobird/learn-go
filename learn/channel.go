package main

import "fmt"

func speak(arg string, ch chan string) {
	ch <- arg // Send
}

func main() {
	ch := make(chan string, 2)

	go speak("Hello World", ch)
	go speak("Hi again", ch)

	data1 := <-ch
	fmt.Println(data1)

	data2, ok := <-ch
	fmt.Println(data2, ok)

	close(ch)

	data3, ok := <-ch
	fmt.Printf("%T, %v\n", data3, ok)

	data4, ok := <-ch
	fmt.Printf("%T, %v\n", data4, ok)

}
