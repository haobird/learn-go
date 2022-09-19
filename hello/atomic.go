package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Person struct {
	Name string
	Age  int32
}

func add(w *sync.WaitGroup, num *int32) {
	defer w.Done()
	atomic.AddInt32(num, 1)
	atomic.Value
}

func main() {
	var n int32 = 0
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i = i + 1 {
		go add(&wg, &n)
	}

	wg.Wait()

	var p = Person{"hh", 12}

	fmt.Println("Result:", n)
}
