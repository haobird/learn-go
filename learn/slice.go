package main

import "fmt"

func main() {
	//
	var arr []string
	// var s []string
	var s = make([]string, 0, 0)

	fmt.Println(s)
	fmt.Println(s == nil)

	for i := 0; i < 10; i++ {
		arr = append(arr, fmt.Sprintf("%d", i))
	}
	fmt.Println(arr)
}
