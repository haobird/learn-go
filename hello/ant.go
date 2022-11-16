package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/panjf2000/ants/v2"
)

var (
	ch       = make(chan string, 10)
	runTimes = 1000
)

func forward(str string, ch chan string) {
	fmt.Printf("%s pushing\n", str)
	ch <- str
	fmt.Printf("%s pushed\n", str)
}

func myFunc(i interface{}) {
	time.Sleep(2 * time.Second)
	fmt.Printf("run with %v\n", i)
}

func main() {

	// 循环写入通道
	go func() {
		for i := 0; i < runTimes; i++ {
			str := fmt.Sprintf("gorouting--1-%d", i)
			forward(str, ch)
		}
	}()

	go func() {
		for i := 0; i < runTimes; i++ {
			str := fmt.Sprintf("gorouting--2-%d", i)
			forward(str, ch)
		}
	}()

	p, _ := ants.NewPoolWithFunc(2, func(i interface{}) {
		myFunc(i)
	})
	defer p.Release()

	// 持续消费通道
	for i := range ch {
		fmt.Printf("[constumer]--%v\n", i)
		_ = p.Invoke(i)
	}

	errs := make(chan error, 2)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	log.Println(fmt.Sprintf(" service terminated: %s", err))

}
