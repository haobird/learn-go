package main

import "time"

func main() {
	WillPanic()
	time.Sleep(5 * time.Minute)
}

func WillPanic() {
	panic("Woah")
}
