package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m = make(map[int]string)

	wg.Add(10)
	for i := 0; i <= 4; i++ {
		go func(k int) {
			v := fmt.Sprintf("value %v", k)

			fmt.Println("Writing:", v)
			m[k] = v
			wg.Done()
		}(i)
	}

	for i := 0; i <= 4; i++ {
		go func(k int) {
			v, _ := m[k]
			fmt.Println("Reading: ", v)
			wg.Done()
		}(i)
	}

	wg.Wait()

	var pr = func(k any, v any) bool {
		fmt.Printf("key %d is %s\n", k, v)
		return true
	}

	for k, v := range m {
		pr(k, v)
	}
}
