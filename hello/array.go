package main

import "fmt"

func main() {
	var a = [4]int{1, 2, 3, 4}
	var b [2]int = a // Error, cannot use a (type [4]int) as type [2]int in assignment

	fmt.Println(b)
}
