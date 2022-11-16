package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/gorm/logger"
)

func main() {
	errs := make(chan error, 2)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	logger.Error(fmt.Sprintf(" service terminated: %s", err))
}
