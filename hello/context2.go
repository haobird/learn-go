package main

import (
	"fmt"
	"sync"
)

func main() {
	// ctx := context.Background()

	var closedchan = make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		val1 := <-closedchan
		fmt.Println("ddd1:", val1)
		wg.Done()
	}()

	go func() {
		val1 := <-closedchan
		fmt.Println("ddd2:", val1)
		wg.Done()
	}()

	close(closedchan)

	wg.Wait()

}
