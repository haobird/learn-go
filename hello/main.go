package main

import "fmt"

var name string = "My name is Go"

var bio string = ``

func main() {
	add := myFunction2()
	add(5)
	fmt.Println(add(10))
	// myFunction3()

	add2 := myFunction2()
	fmt.Println(add2(10))

	sum := myFunction2()(10)
	fmt.Println(sum)
	fmt.Println(myFunction2()(10))
}

func myFunction1() {
	fn := func(str string) {
		fmt.Println("inside fn", str)
	}

	fn("kk")
}

func myFunction2() func(int) int {
	sum := 0

	return func(v int) int {
		fmt.Println(sum)
		sum += v
		return sum
	}
}

func myFunction3() {
	sum := 0

	fn := func(v int) int {
		fmt.Println(sum)
		sum += v
		return sum
	}

	fn(10)
	fmt.Println(fn(5))
}
