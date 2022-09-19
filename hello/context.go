package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handler started")

	context := req.Context()

	go func() {
		val := <-context.Done()
		fmt.Println("val:", val)
		fmt.Println(context.Err())
	}()

	select {
	// Simulating some work by the server, waits 5 seconds and then responds.
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Response from the server")

	// Handling request cancellation
	case val1 := <-context.Done():
		err := context.Err()
		fmt.Println("Error:", err, val1)
	}

	fmt.Println("Handler complete")
}

func main() {
	http.HandleFunc("/request", handleRequest)

	fmt.Println("Server is running...")
	http.ListenAndServe(":4000", nil)
}
