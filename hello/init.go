package main

import (
	"fmt"
	"hello/pkg"
)

func init() {
	fmt.Println("Before main!1")
}

func init() {
	fmt.Println("Hello again?1")
}

func main() {
	fmt.Println("Running main")
	pkg.Print()
}
