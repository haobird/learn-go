package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctxx := context.WithValue(ctx, "x", "1")
	ctxxx := context.WithValue(ctxx, "y", "2")
	ctxxxx := context.WithValue(ctxxx, "x", "3")

	fmt.Println(ctx.Value("x"))    // nil
	fmt.Println(ctxx.Value("x"))   // 1
	fmt.Println(ctxxx.Value("x"))  // 1
	fmt.Println(ctxxxx.Value("x")) // 3
	fmt.Println(ctxxxx.Value("y")) // 2
}
