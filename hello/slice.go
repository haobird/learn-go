package main

import "fmt"

func main() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := s1

	s2[0] = "Sun"

	fmt.Println(s1) // Output: [Sun Tue Wed Thu Fri Sat Sun]
	fmt.Println(s2) // Output: [Sun Tue]
}
