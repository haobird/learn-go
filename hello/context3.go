package main

import (
	"fmt"
	"time"

	"golang.org/x/example/context"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)

	ctx1, _ := context.WithCancel(ctx)
	fmt.Println(ctx1)

	ctx2, cancel2 := context.WithTimeout(ctx1, time.Minute)
	fmt.Println(ctx2)

	fmt.Println(ctx2.Err())

	cancel2()
	fmt.Println(ctx2.Err())

}
