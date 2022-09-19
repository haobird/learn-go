package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(10)
	for i := 0; i <= 4; i++ {
		go func(k int) {
			v := fmt.Sprintf("value %v", k)

			fmt.Println("Writing:", v)
			m.Store(k, v)
			wg.Done()
		}(i)
	}

	for i := 0; i <= 4; i++ {
		go func(k int) {
			v, _ := m.Load(k)
			fmt.Println("Reading: ", v)
			wg.Done()
		}(i)
	}

	wg.Wait()

	var pr = func(k any, v any) bool {
		fmt.Printf("key %d is %s\n", k, v)
		return true
	}

	m.Range(pr)
}
