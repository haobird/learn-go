package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	increment := func() {
		count++
	}

	increment2 := func() {
		count += 2
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)

	once.Do(increment2)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
