package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"
)

type connRequest struct {
	conn string
	err  error
}

func main() {
	// ctx := context.Background()
	// req := make(chan connRequest, 1)
	// // waitStart := nowFunc()

	// go func() {
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("ddd")
	// 	case ret, ok := <-req:
	// 		fmt.Println("xxx")
	// 		fmt.Println(ret, ok)
	// 	}
	// }()

	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	log.Panic(fmt.Sprintf(" service terminated: %s", err))
}

var nowFunc = time.Now
