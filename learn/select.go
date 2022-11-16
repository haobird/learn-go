package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	close(one)
	close(two)
	go print()
	select {}

}

func print() {
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}

}
