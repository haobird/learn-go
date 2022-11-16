package main

import (
	"fmt"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	fmt.Println(name, "reading")
	for !done {
		fmt.Println(name, "reading dddd")
		c.Wait()
		fmt.Println(name, "reading dddd222")
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func read1(name string, c *sync.Cond) {
	c.L.Lock()
	fmt.Println(name, "reading")
	c.Wait()
	fmt.Println(name, "starts reading")
	// c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second)

	// c.L.Lock()
	done = true
	fmt.Println("ddd")
	// c.L.Unlock()

	fmt.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	var m sync.Mutex
	cond := sync.NewCond(&m)

	go read1("Reader 1", cond)
	go read1("Reader 2", cond)
	go read1("Reader 3", cond)
	write("Writer", cond)

	time.Sleep(4 * time.Second)
}
