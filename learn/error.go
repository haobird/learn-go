package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println(err)
		// Do something with the error
		default:
			fmt.Println("no idea!")
		}

		return
	}

	fmt.Println(result)
}

func Divide1(a, b int) int {
	return a / b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func Divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a / b, nil
}

func Divide3(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a / b, nil
}

func Divide4()
