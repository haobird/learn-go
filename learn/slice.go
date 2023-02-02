package main

import "fmt"

func main() {
	str := "C7C00600FF0403110727414B"
	fmt.Println(len(str))
	n := 18
	fmt.Println(str[n:n+5] + "/" + str[n+5:])
	// var arr []string
	// var s []string

	// fmt.Println(s)
	// fmt.Println(s == nil)
	// s = append(s, "a")
	// fmt.Println(s)

	// var arr = make([]string, 0, 0)
	// fmt.Println(arr)
	// for i := 0; i < 10; i++ {
	// 	arr = append(arr, fmt.Sprintf("%d", i))
	// }
	// fmt.Println(arr)
}
