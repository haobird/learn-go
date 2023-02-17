package main

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

var (
	ErrTwoZeroes   = errors.New("can't sum two zeroes")
	ErrIntOverflow = errors.New("integer overflow")
)

type ErrorWrapper struct {
	ErrCode int32  `json:"errorCode"`
	ErrMsg  string `json:"errorMsg"`
}

func (e ErrorWrapper) Error() string {
	return e.ErrMsg
}

func Sum(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, &ErrorWrapper{-1, ErrTwoZeroes.Error()}
	}
	if (b > 0 && a > (intMax-b)) || (b < 0 && a < (intMin-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

func main() {
	sum, err := Sum(0, 1)
	fmt.Println(sum, err)

	sum, err = Sum(0, 0)
	if err != nil {
		fmt.Println("type:", reflect.TypeOf(err))
		var errorWrapper *ErrorWrapper
		if errors.As(err, &errorWrapper) {
			fmt.Println(errorWrapper.ErrCode, errorWrapper.ErrMsg)
		}
		fmt.Println(err)
	}

}
