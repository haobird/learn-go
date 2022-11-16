package main

import (
	"context"
	"fmt"
	"time"
)

var (
	processID = "abc-xyz"
)

func main() {
	ctx := context.Background()
	// fmt.Println(ctx.Done())
	var ch chan struct{}

	val := <-ch
	fmt.Println(val)

	select {
	// Simulating some work by the server, waits 5 seconds and then responds.
	case <-time.After(5 * time.Second):
		fmt.Println("Response from the server")

	// Handling request cancellation
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Error:", err)
	}

	fmt.Println("Handler complete")

	// time.Sleep(10 * time.Minute)
	// ctx = context.WithValue(ctx, "processID", processID)

	// ProcessRequest(ctx)
}

func ProcessRequest(ctx context.Context) {
	value := ctx.Value("processID")
	fmt.Printf("Processing ID: %v", value)
}
